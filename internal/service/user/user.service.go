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
	"github.com/rs/zerolog/log"
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
	tokenId := request.GetToken()
	userId := request.GetUserId()

	var token token_ent.Token
	err := s.tokenRepo.GetTokenInfo(tokenId, &token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info().Err(err).
				Str("service", "checkin").
				Str("module", "AddEvent").
				Str("token", tokenId).
				Msg("someone is trying to use invalid token")

			return nil, status.Error(codes.NotFound, "Token not found")
		}

		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEvent").
			Str("token", tokenId).
			Msg("Error while getting token infomation")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	if token.EndAt < time.Now().Unix() {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEvent").
			Str("token", tokenId).
			Int64("end_at", token.EndAt).
			Msg("Token expired")

		return nil, status.Error(codes.DeadlineExceeded, "Invalid token")
	}

	eventId := token.Event.EventId

	err = s.userRepo.IsEventTaken(userId, eventId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEvent").
			Str("user_id", userId).
			Str("event_id", eventId).
			Msg("Error while checking if user_id has already taken event_id")

		return nil, status.Error(codes.Internal, "Internal server error")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEvent").
			Str("user_id", userId).
			Str("event_id", eventId).
			Msg("user_id has already taken event_id")

		return nil, status.Error(codes.AlreadyExists, "User has already taken the event")
	}

	takenAt := time.Now().Unix()
	userEvent := event_ent.UserEvent{
		UserId:  userId,
		EventId: eventId,
		TakenAt: takenAt,
	}

	err = s.userRepo.AddEvent(userEvent)
	if err != nil {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEvent").
			Str("user_id", userId).
			Str("event_id", eventId).
			Int64("taken_at", takenAt).
			Msg("Error while appending user event to database")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	proto := token.Event.ToProto()

	res := &v1.AddEventResponse{
		Event: proto,
	}

	return res, nil
}

func (s *UserService) GetAllUserEventsByNamespaceId(ctx context.Context, request *v1.GetAllUserEventsByNamespaceIdRequest) (*v1.GetAllUserEventsByNamespaceIdResponse, error) {
	userId := request.GetUserId()
	namespaceId := request.GetNamespaceId()

	var userEvents []*event_ent.UserEvent

	err := s.userRepo.GetUserEventsByNamespaceId(userId, namespaceId, &userEvents)
	if err != nil {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GetAllUserEventsByNamespaceId").
			Str("user_id", userId).
			Str("namespace_id", namespaceId).
			Msg("Error while getting all events of namespace_id in user_id")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	var events []*event_proto.UserEvent
	for _, userEvent := range userEvents {
		if userEvent.Event.NamespaceId != namespaceId {
			continue
		}

		events = append(events, &event_proto.UserEvent{
			Event:   userEvent.Event.ToProto(),
			IsTaken: true,
			TakenAt: userEvent.TakenAt,
		})
	}

	return &v1.GetAllUserEventsByNamespaceIdResponse{
		Event: events,
	}, nil
}

func (s *UserService) GetUserEventByEventId(ctx context.Context, request *v1.GetUserEventByEventIdRequest) (*v1.GetUserEventByEventIdResponse, error) {
	userId := request.GetUserId()
	eventId := request.GetEventId()

	var userEvent event_ent.UserEvent

	err := s.userRepo.GetUserEventById(userId, eventId, &userEvent)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).
				Str("service", "checkin").
				Str("module", "GetUserEventByEventId").
				Str("user_id", userId).
				Str("event_id", eventId).
				Msg("user_id has never taken event_id")

			return nil, status.Error(codes.NotFound, "User has never taken the event")
		} else {
			log.Error().Err(err).
				Str("service", "checkin").
				Str("module", "GetUserEventByEventId").
				Str("user_id", userId).
				Str("event_id", eventId).
				Msg("Error while getting event_id of user_id")

			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	return &v1.GetUserEventByEventIdResponse{
		UserEvent: &event_proto.UserEvent{
			Event:   userEvent.Event.ToProto(),
			IsTaken: true,
			TakenAt: userEvent.TakenAt,
		},
	}, nil
}
