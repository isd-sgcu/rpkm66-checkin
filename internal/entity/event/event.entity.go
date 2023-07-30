package event

import (
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/event/v1"
)

type Event struct {
	EventId        string `gorm:"column:event_id;unique;not null;primaryKey"`
	EventName      string `gorm:"column:event_name;not null"`
	NamespaceId    string `gorm:"column:namespace_id;not null"`
	AdditionalInfo string `gorm:"column:additional_info"`
}

func (e *Event) ToProto() *v1.Event {
	return &v1.Event{
		EventId:        e.EventId,
		EventName:      e.EventName,
		NamespaceId:    e.NamespaceId,
		AdditionalInfo: e.AdditionalInfo,
	}
}
