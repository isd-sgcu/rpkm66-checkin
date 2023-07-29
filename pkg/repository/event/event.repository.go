package event

import (
	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/internal/repository/event"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllEvents(result *[]*event_ent.Event) error
	GetEventByEventId(eventId string, result *event_ent.Event) error
	GetEventsByUserId(userId string, result *[]*event_ent.UserEvent) error
	GetEventsByNamespaceId(namespaceId string, result *[]*event_ent.Event) error
	GetEventByToken(token string, result *event_ent.Event) error
}

func NewRepository(db *gorm.DB) Repository {
	return event_repo.NewRepository(db)
}
