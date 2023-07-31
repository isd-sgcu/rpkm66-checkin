package event

import (
	"context"
	"errors"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/event/v1"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GetAllEvents").
			Msg("Error while getting all events")

		return nil, status.Error(codes.Internal, "Internal server error")
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
	eventId := request.GetEventId()

	var event event_ent.Event
	err := s.repo.GetEventByEventId(eventId, &event)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).
				Str("service", "checkin").
				Str("module", "GetEventByEventId").
				Str("event_id", eventId).
				Msg("Cannot find event with the specified event_id")

			return nil, status.Error(codes.NotFound, "Event not found")
		} else {
			log.Error().Err(err).
				Str("service", "checkin").
				Str("module", "GetEventByEventId").
				Str("event_id", eventId).
				Msg("Error while getting all events")

			return nil, status.Error(codes.Internal, "Internal server error")
		}
	}

	var proto *v1.Event = event.ToProto()

	res := &v1.GetEventByEventIdResponse{
		Event: proto,
	}

	return res, nil
}

func (s *EventService) GetEventsByUserId(ctx context.Context, request *v1.GetEventsByUserIdRequest) (*v1.GetEventsByUserIdResponse, error) {
	userId := request.GetUserId()

	var userEvents []*event_ent.UserEvent
	// No way of getting Record Not Found
	err := s.repo.GetEventsByUserId(userId, &userEvents)
	if err != nil {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GetEventsByUserId").
			Str("user_id", userId).
			Msg("Error while getting all events of user_id")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	userEventMap := make(map[string]*event_ent.UserEvent)
	for _, userEvent := range userEvents {
		userEventMap[userEvent.EventId] = userEvent
	}

	var events []*event_ent.Event
	err = s.repo.GetAllEvents(&events)
	if err != nil {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GetEventsByUserId").
			Msg("Error while getting all events")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	protos := make([]*v1.UserEvent, len(events))
	for i, event := range events {
		userEvent, ok := userEventMap[event.EventId]
		var takenAt int64 = -1
		if ok {
			takenAt = userEvent.TakenAt
		}

		protos[i] = &v1.UserEvent{
			Event:   event.ToProto(),
			IsTaken: ok,
			TakenAt: takenAt,
		}
	}

	res := &v1.GetEventsByUserIdResponse{
		Events: protos,
	}

	return res, nil
}

func (s *EventService) GetEventsByNamespaceId(ctx context.Context, request *v1.GetEventsByNamespaceIdRequest) (*v1.GetEventsByNamespaceIdResponse, error) {
	namespaceId := request.GetNamespaceId()

	var events []*event_ent.Event
	err := s.repo.GetEventsByNamespaceId(namespaceId, &events)
	if err != nil {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GetEventsByNamespaceId").
			Str("namespace_id", namespaceId).
			Msg("Error while getting all events of namespace_id")

		return nil, status.Error(codes.Internal, "Internal server error")
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
