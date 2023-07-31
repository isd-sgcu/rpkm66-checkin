package staff

import (
	"context"
	"errors"
	"time"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/staff/v1"
	"github.com/isd-sgcu/rpkm66-checkin/internal/utils"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	staff_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/staff"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type StaffService struct {
	v1.UnimplementedStaffServiceServer
	staffRepo staff_repo.Repository
	userRepo  user_repo.Repository
	eventRepo event_repo.Repository
}

func NewService(staffRepo staff_repo.Repository, userRepo user_repo.Repository, eventRepo event_repo.Repository) v1.StaffServiceServer {
	return &StaffService{
		v1.UnimplementedStaffServiceServer{},
		staffRepo,
		userRepo,
		eventRepo,
	}
}

func (s *StaffService) IsStaff(ctx context.Context, request *v1.IsStaffRequest) (*v1.IsStaffResponse, error) {
	staffId := request.GetStaffId()

	isStaff, err := s.staffRepo.IsStaff(staffId)
	// Since record not found means staff_id is not a staff
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "IsStaff").
			Str("staff_id", staffId).
			Msg("Error while checking if staff_id is staff")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	res := &v1.IsStaffResponse{
		IsStaff: isStaff,
	}

	return res, nil
}

func (s *StaffService) AddEventToUser(ctx context.Context, request *v1.AddEventToUserRequest) (*v1.AddEventToUserResponse, error) {
	staffUserId := request.GetStaffUserId()
	eventId := request.GetEventId()
	userId := request.GetUserId()

	isStaff, err := s.staffRepo.IsStaff(staffUserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEventToUser").
			Str("user_id", staffUserId).
			Msg("Error while checking if user_id is staff")

		return nil, status.Error(codes.Internal, "Internal server error")
	} else if !isStaff {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEventToUser").
			Str("staff_id", staffUserId).
			Msg("staff_id is not registered as staff")

		return nil, status.Error(codes.PermissionDenied, "Only staff user can perform this action.")
	}

	doesExist, err := s.eventRepo.DoesEventExist(eventId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEventToUser").
			Str("event_id", eventId).
			Msg("Error while checking if event_id exist")

		return nil, status.Error(codes.Internal, "Internal server error")
	} else if !doesExist {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEventToUser").
			Str("event_id", eventId).
			Msg("event_id does not exist")

		return nil, status.Error(codes.InvalidArgument, "Invalid event id")
	}

	isTaken, err := s.userRepo.IsEventTaken(userId, eventId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEventToUser").
			Str("user_id", userId).
			Str("event_id", eventId).
			Msg("Error while checking if user_id has already taken event_id")

		return nil, status.Error(codes.Internal, "Internal server error")
	} else if isTaken {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEventToUser").
			Str("user_id", userId).
			Str("event_id", eventId).
			Msg("user_id has already taken event_id")

		return nil, status.Error(codes.AlreadyExists, "User has already taken the event")
	}

	takenAt := time.Now().Unix()
	userEvent := event_ent.UserEvent{
		UserId:  userId,
		EventId: eventId,
		TakenAt: takenAt,
	}

	err = s.userRepo.AddEvent(userEvent)
	if err != nil {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "AddEventToUser").
			Str("user_id", userId).
			Str("event_id", eventId).
			Int64("taken_at", takenAt).
			Msg("Error while appending user event to database")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	res := &v1.AddEventToUserResponse{
		Success: true,
	}

	return res, nil
}

func (s *StaffService) GenerateSignInToken(ctx context.Context, request *v1.GenerateSignInTokenRequest) (*v1.GenerateSignInTokenResponse, error) {
	staffUserId := request.GetStaffUserId()
	eventId := request.GetEventId()

	isStaff, err := s.staffRepo.IsStaff(staffUserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GenerateSignInToken").
			Str("staff_id", staffUserId).
			Msg("Error while checking if staff_id is staff")

		return nil, status.Error(codes.Internal, "Internal server error")
	} else if !isStaff {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GenerateSignInToken").
			Str("staff_id", staffUserId).
			Msg("staff_id is not registered as staff")

		return nil, status.Error(codes.PermissionDenied, "Only staff user can perform this action.")
	}

	doesExist, err := s.eventRepo.DoesEventExist(eventId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GenerateSignInToken").
			Str("event_id", eventId).
			Msg("Error while checking if event_id exist")

		return nil, status.Error(codes.Internal, "Internal server error")
	} else if !doesExist {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GenerateSignInToken").
			Str("event_id", eventId).
			Msg("event_id does not exist")

		return nil, status.Error(codes.InvalidArgument, "Invalid event id")
	}

	token := utils.GenToken(eventId + staffUserId)
	endAt := time.Now().Add(time.Minute * 15).Unix() // token valids for 15 minutes
	tokenEntity := token_ent.Token{
		Id:      token,
		EventId: eventId,
		EndAt:   endAt,
	}

	err = s.staffRepo.CreateToken(tokenEntity)
	if err != nil {
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GenerateSignInToken").
			Str("token", token).
			Str("event_id", eventId).
			Int64("end_at", endAt).
			Msg("Error while appending token data to database")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	res := &v1.GenerateSignInTokenResponse{
		Token: token,
	}

	return res, nil
}
