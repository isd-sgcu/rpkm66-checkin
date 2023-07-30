package token

import "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"

type Token struct {
	Id      string      `gorm:"column:id;unique;not null;primaryKey"`
	EventId string      `gorm:"column:event_id;not null"`
	Event   event.Event `gorm:"references:event_id"`
	EndAt   int64       `gorm:"column:end_at;not null"`
}
