package namespace

import v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/namespace/v1"

type Namespace struct {
	id string `gorm:"column:id"`
}

func (e *Namespace) ToProto() *v1.Namespace {
	return &v1.Namespace{
		Id: e.id,
	}
}
