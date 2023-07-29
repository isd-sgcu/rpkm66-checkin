package staff

import (
	"context"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/user/v1"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	user_repo  user.Repository
	event_repo event.Repository
}

func NewService(user_repo user.Repository, event_repo event.Repository) v1.UserServiceServer {
	return &UserService{
		v1.UnimplementedUserServiceServer{},
		user_repo,
		event_repo,
	}
}

func (s *UserService) AddEvent(ctx context.Context, request *v1.AddEventRequest) (*v1.AddEventResponse, error) {
	var event event_ent.Event
	err := s.event_repo.GetEventByToken(request.GetToken(), &event)
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
