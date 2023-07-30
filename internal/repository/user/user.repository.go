package user

import (
	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) AddEvent(userEvent event_ent.UserEvent) error {
	return r.db.Create(userEvent).Error
}
