package staff

import (
	"context"
	"errors"
	"time"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/user/v1"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	token_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/token"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	user_repo  user_repo.Repository
	event_repo event_repo.Repository
	token_repo token_repo.Repository
}

func NewService(user_repo user_repo.Repository, event_repo event_repo.Repository, token_repo token_repo.Repository) v1.UserServiceServer {
	return &UserService{
		v1.UnimplementedUserServiceServer{},
		user_repo,
		event_repo,
		token_repo,
	}
}

func (s *UserService) AddEvent(ctx context.Context, request *v1.AddEventRequest) (*v1.AddEventResponse, error) {
	var token token_ent.Token
	err := s.token_repo.GetTokenInfo(request.GetToken(), &token)
	if err != nil {
		return nil, err
	}

	if token.EndAt < time.Now().Unix() {
		return nil, errors.New("Invalid token")
	}

	var event event_ent.Event
	err = s.event_repo.GetEventByEventId(token.EventId, &event)
	if err != nil {
		return nil, err
	}

	err = s.user_repo.AddEvent(request.UserId, event.EventId)
	if err != nil {
		return nil, err
	}

	proto := event.ToProto()

	res := &v1.AddEventResponse{
		Event: proto,
	}

	return res, nil
}
