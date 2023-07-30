package staff

import (
	"context"
	"errors"
	"time"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	event_proto "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/event/v1"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/user/v1"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	token_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/token"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	userRepo  user_repo.Repository
	eventRepo event_repo.Repository
	tokenRepo token_repo.Repository
}

func NewService(userRepo user_repo.Repository, eventRepo event_repo.Repository, tokenRepo token_repo.Repository) v1.UserServiceServer {
	return &UserService{
		v1.UnimplementedUserServiceServer{},
		userRepo,
		eventRepo,
		tokenRepo,
	}
}

func (s *UserService) AddEvent(ctx context.Context, request *v1.AddEventRequest) (*v1.AddEventResponse, error) {
	var token token_ent.Token
	err := s.tokenRepo.GetTokenInfo(request.GetToken(), &token)
	if err != nil {
		return nil, err
	}

	if token.EndAt < time.Now().Unix() {
		return nil, status.Error(codes.DeadlineExceeded, "Invalid token")
	}

	var event event_ent.Event
	err = s.eventRepo.GetEventByEventId(token.EventId, &event)
	if err != nil {
		return nil, err
	}

	isTaken, err := s.userRepo.IsEventTaken(request.GetUserId(), event.EventId)
	if err != nil {
		return nil, err
	}

	if isTaken {
		return nil, status.Error(codes.AlreadyExists, "Event has already been taken")
	}

	userEvent := event_ent.UserEvent{
		UserId:  request.GetUserId(),
		EventId: event.EventId,
		TakenAt: time.Now().Unix(),
	}

	err = s.userRepo.AddEvent(userEvent)
	if err != nil {
		return nil, err
	}

	proto := event.ToProto()

	res := &v1.AddEventResponse{
		Event: proto,
	}

	return res, nil
}

func (s *UserService) GetAllUserEventsByNamespaceId(ctx context.Context, request *v1.GetAllUserEventsByNamespaceIdRequest) (*v1.GetAllUserEventsByNamespaceIdResponse, error) {
	userId := request.GetUserId()
	namespaceId := request.GetNamespaceId()

	var userEvents []*event_ent.UserEvent

	err := s.eventRepo.GetEventsByUserId(userId, &userEvents)
	if err != nil {
		// TODO: LOG
		// No need to check for not found because it is not possible
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	eventsId := make([]*string, len(userEvents))
	takenAtMap := make(map[string]int64, 0)
	for _, userEvent := range userEvents {
		eventsId = append(eventsId, &userEvent.EventId)
		takenAtMap[userEvent.EventId] = userEvent.TakenAt
	}

	var events []*event_ent.Event
	err = s.eventRepo.GetEventByEventIdsWithNamespace(&eventsId, namespaceId, &events)
	if err != nil {
		// TODO: LOG
		// No need to check for not found because it is not possible
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	var result []*event_proto.UserEvent

	for _, event := range events {
		result = append(result, &event_proto.UserEvent{
			Event:   event.ToProto(),
			IsTaken: true,
			TakenAt: takenAtMap[event.EventId],
		})
	}

	return &v1.GetAllUserEventsByNamespaceIdResponse{
		Event: result,
	}, nil
}

func (s *UserService) GetUserEventByEventId(ctx context.Context, request *v1.GetUserEventByEventIdRequest) (*v1.GetUserEventByEventIdResponse, error) {
	userId := request.GetUserId()
	eventId := request.GetEventId()

	var userEvent *event_ent.UserEvent

	err := s.userRepo.GetUserEventById(userId, eventId, userEvent)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "User does not have this event")
		}
		// TODO: LOG
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	var event *event_ent.Event
	err = s.eventRepo.GetEventByEventId(eventId, event)
	if err != nil {
		// TODO: LOG
		// should not be possible due to event-user_events foreign key constraint
		return nil, status.Error(codes.Internal, "Cannot find given event")
	}

	return &v1.GetUserEventByEventIdResponse{
		UserEvent: &event_proto.UserEvent{
			Event:   event.ToProto(),
			IsTaken: true,
			TakenAt: userEvent.TakenAt,
		},
	}, nil
}
