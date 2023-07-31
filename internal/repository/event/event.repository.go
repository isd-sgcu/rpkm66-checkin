package event

import (
	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		db,
	}
}

func (r *EventRepository) GetAllEvents(result *[]*event_ent.Event) error {
	return r.db.Model(&event_ent.Event{}).Find(result).Error
}

func (r *EventRepository) GetEventByEventId(eventId string, result *event_ent.Event) error {
	return r.db.First(result, "event_id = ?", eventId).Error
}

func (r *EventRepository) GetEventsByUserId(userId string, result *[]*event_ent.UserEvent) error {
	return r.db.Model(&event_ent.UserEvent{}).Find(result, "user_id = ?", userId).Error
}

func (r *EventRepository) GetEventsByNamespaceId(namespaceId string, result *[]*event_ent.Event) error {
	return r.db.Find(result, "namespace_id = ?", namespaceId).Error
}

func (r *EventRepository) DoesEventExist(eventId string) (bool, error) {
	var event event_ent.Event
	err := r.db.First(&event, "event_id = ?", eventId).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
