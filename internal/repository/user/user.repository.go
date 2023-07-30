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

func (r *UserRepository) IsEventTaken(userId string, eventId string) (bool, error) {
	var result event_ent.UserEvent
	err := r.db.Model(&event_ent.UserEvent{}).Find(&result, "user_id = ? AND event_id = ?", userId, eventId).Error
	if err != nil {
		return false, err
	}

	isTaken := result.EventId != "" || result.UserId != ""

	return isTaken, err
}
