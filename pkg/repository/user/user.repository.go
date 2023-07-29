package user

import (
	user_repo "github.com/isd-sgcu/rpkm66-checkin/internal/repository/user"
	"gorm.io/gorm"
)

type Repository interface {
	AddEvent(user_id string, event_id string) error
}

func NewRepository(db *gorm.DB) Repository {
	return user_repo.NewRepository(db)
}
