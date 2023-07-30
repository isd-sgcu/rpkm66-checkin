package user

import (
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/user/v1"
	user_service "github.com/isd-sgcu/rpkm66-checkin/internal/service/user"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	token_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/token"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
)

func NewService(user_repo user_repo.Repository, event_repo event_repo.Repository, token_repo token_repo.Repository) v1.UserServiceServer {
	return user_service.NewService(user_repo, event_repo, token_repo)
}
