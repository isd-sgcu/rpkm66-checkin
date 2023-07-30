package user

import (
	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/internal/repository/user"
	"gorm.io/gorm"
)

type Repository interface {
	AddEvent(userEvent event_ent.UserEvent) error
	IsEventTaken(userId string, eventId string) (bool, error)
	GetUserEventById(userId string, eventId string, userEvent *event_ent.UserEvent) error
	GetUserEventsByNamespaceId(userId string, namespaceId string, userEvent *[]*event_ent.UserEvent) error
}

func NewRepository(db *gorm.DB) Repository {
	return user_repo.NewRepository(db)
}
