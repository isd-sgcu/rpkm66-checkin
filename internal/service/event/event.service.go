package event

import (
	"context"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/event/v1"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
)

type EventService struct {
	v1.UnimplementedEventServiceServer
	repo event.Repository
}

func NewService(repo event.Repository) v1.EventServiceServer {
	return &EventService{
		v1.UnimplementedEventServiceServer{},
		repo,
	}
}

func (s *EventService) GetAllEvents(ctx context.Context, request *v1.GetAllEventsRequest) (*v1.GetAllEventsResponse, error) {
	var events []*event_ent.Event
	err := s.repo.GetAllEvents(&events)
	if err != nil {
		return nil, err
	}

	protos := make([]*v1.Event, len(events))
	for i, event := range events {
		protos[i] = event.ToProto()
	}

	res := &v1.GetAllEventsResponse{
		Events: protos,
	}

	return res, nil
}

func (s *EventService) GetEventByEventId(ctx context.Context, request *v1.GetEventByEventIdRequest) (*v1.GetEventByEventIdResponse, error) {
	var event event_ent.Event
	err := s.repo.GetEventByEventId(request.GetEventId(), &event)
	if err != nil {
		return nil, err
	}

	var proto *v1.Event = event.ToProto()

	res := &v1.GetEventByEventIdResponse{
		Event: proto,
	}

	return res, nil
}

func (s *EventService) GetEventsByUserId(ctx context.Context, request *v1.GetEventsByUserIdRequest) (*v1.GetEventsByUserIdResponse, error) {
	var userEvents []*event_ent.UserEvent
	err := s.repo.GetEventsByUserId(request.GetUserId(), &userEvents)
	if err != nil {
		return nil, err
	}

	// create a set of event
	userEventMap := make(map[string]*event_ent.UserEvent)
	for _, userEvent := range userEvents {
		userEventMap[userEvent.EventId] = userEvent
	}

	var events []*event_ent.Event
	err = s.repo.GetAllEvents(&events)
	if err != nil {
		return nil, err
	}

	protos := make([]*v1.UserEvent, len(events))
	for i, event := range events {
		userEvent, ok := userEventMap[event.EventId]
		var when int64 = -1
		if ok {
			when = userEvent.When
		}

		protos[i] = &v1.UserEvent{
			Event:   event.ToProto(),
			IsTaken: ok,
			When:    when,
		}
	}

	res := &v1.GetEventsByUserIdResponse{
		Events: protos,
	}

	return res, nil
}

func (s *EventService) GetEventsByNamespaceId(ctx context.Context, request *v1.GetEventsByNamespaceIdRequest) (*v1.GetEventsByNamespaceIdResponse, error) {
	var events []*event_ent.Event
	err := s.repo.GetEventsByNamespaceId(request.GetNamespaceId(), &events)
	if err != nil {
		return nil, err
	}

	protos := make([]*v1.Event, len(events))
	for i, event := range events {
		protos[i] = event.ToProto()
	}

	res := &v1.GetEventsByNamespaceIdResponse{
		Events: protos,
	}

	return res, nil
}
