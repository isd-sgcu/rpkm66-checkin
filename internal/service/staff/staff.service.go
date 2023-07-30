package staff

import (
	"context"
	"time"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/staff/v1"
	"github.com/isd-sgcu/rpkm66-checkin/internal/utils"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	staff_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/staff"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	isStaff, err := s.staffRepo.IsStaff(request.GetStaffId())
	if err != nil {
		return nil, err
	}

	res := &v1.IsStaffResponse{
		IsStaff: isStaff,
	}

	return res, nil
}

func (s *StaffService) AddEventToUser(ctx context.Context, request *v1.AddEventToUserRequest) (*v1.AddEventToUserResponse, error) {
	isStaff, err := s.staffRepo.IsStaff(request.GetStaffUserId())
	if err != nil {
		return nil, err
	}

	if !isStaff {
		return nil, status.Error(codes.PermissionDenied, "Only staff user can perform this action.")
	}

	ok, err := s.eventRepo.DoesEventExist(request.GetEventId())
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Invalid event id")
	}

	isTaken, err := s.userRepo.IsEventTaken(request.GetUserId(), request.GetEventId())
	if err != nil {
		return nil, err
	}

	if isTaken {
		return nil, status.Error(codes.AlreadyExists, "User has already taken the event")
	}

	userEvent := event_ent.UserEvent{
		UserId:  request.GetUserId(),
		EventId: request.GetEventId(),
		TakenAt: time.Now().Unix(),
	}

	err = s.userRepo.AddEvent(userEvent)
	if err != nil {
		return nil, err
	}

	res := &v1.AddEventToUserResponse{
		Success: true,
	}

	return res, nil
}

func (s *StaffService) GenerateSignInToken(ctx context.Context, request *v1.GenerateSignInTokenRequest) (*v1.GenerateSignInTokenResponse, error) {
	isStaff, err := s.staffRepo.IsStaff(request.GetStaffUserId())
	if err != nil {
		return nil, err
	}

	if !isStaff {
		return nil, status.Error(codes.PermissionDenied, "Only staff user can perform this action.")
	}

	ok, err := s.eventRepo.DoesEventExist(request.GetEventId())
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Invalid event id")
	}

	token := utils.GenToken(request.GetEventId() + request.GetStaffUserId())
	tokenEntity := token_ent.Token{
		Id:      token,
		EventId: request.GetEventId(),
		EndAt:   time.Now().Add(time.Minute * 15).Unix(), // token valids for 15 minutes
	}

	err = s.staffRepo.CreateToken(tokenEntity)
	if err != nil {
		return nil, err
	}

	res := &v1.GenerateSignInTokenResponse{
		Token: token,
	}

	return res, nil
}
