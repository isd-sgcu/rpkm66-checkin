package namespace

import (
	"github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/namespace/v1"
)

type Namespace struct {
	Id     string `gorm:"column:id;unique;not null"`
	Events []event.Event
}

func (e *Namespace) ToProto() *v1.Namespace {
	return &v1.Namespace{
		Id: e.Id,
	}
}
