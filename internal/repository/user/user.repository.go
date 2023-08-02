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

func (r *UserRepository) IsEventTaken(userId string, eventId string) error {
	var result event_ent.UserEvent
	return r.db.Model(&event_ent.UserEvent{}).First(&result, "user_id = ? AND event_id = ?", userId, eventId).Error
}

func (r *UserRepository) GetUserEventById(userId string, eventId string, userEvent *event_ent.UserEvent) error {
	return r.db.Model(&event_ent.UserEvent{}).Preload("Event").First(userEvent, "user_id = ? AND event_id = ?", userId, eventId).Error
}

func (r *UserRepository) GetUserEventsByNamespaceId(userId string, namespaceId string, userEvents *[]*event_ent.UserEvent) error {
	return r.db.Model(&event_ent.UserEvent{}).Preload("Event").Where("user_id = ?", userId).Find(userEvents).Error
}
