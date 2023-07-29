package user

import (
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/user/v1"
	user_svc "github.com/isd-sgcu/rpkm66-checkin/internal/service/user"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
)

func NewService(user_repo user_repo.Repository, event_repo event_repo.Repository) v1.UserServiceServer {
	return user_svc.NewService(user_repo, event_repo)
}
