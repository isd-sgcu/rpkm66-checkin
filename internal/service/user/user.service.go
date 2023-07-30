package staff

import (
	"context"
	"time"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/user/v1"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	token_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/token"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
