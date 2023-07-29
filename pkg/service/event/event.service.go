package event

import (
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/event/v1"
	event_svc "github.com/isd-sgcu/rpkm66-checkin/internal/service/event"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
)

func NewService(repo event_repo.Repository) v1.EventServiceServer {
	return event_svc.NewService(repo)
}
