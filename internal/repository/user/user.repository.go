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

func (r *UserRepository) AddEvent(user_id string, event_id string) error {
	userEvent := event_ent.UserEvent{
		UserId:  user_id,
		EventId: event_id,
	}

	return r.db.Create(userEvent).Error
}
