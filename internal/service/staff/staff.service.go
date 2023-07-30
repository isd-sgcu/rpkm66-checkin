package staff

import (
	"context"
	"errors"
	"time"

	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/staff/v1"
	"github.com/isd-sgcu/rpkm66-checkin/internal/utils"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/staff"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
)

type StaffService struct {
	v1.UnimplementedStaffServiceServer
	staff_repo staff.Repository
	user_repo  user.Repository
}

func NewService(staff_repo staff.Repository, user_repo user.Repository) v1.StaffServiceServer {
	return &StaffService{
		v1.UnimplementedStaffServiceServer{},
		staff_repo,
		user_repo,
	}
}

func (s *StaffService) IsStaff(ctx context.Context, request *v1.IsStaffRequest) (*v1.IsStaffResponse, error) {
	isStaff, err := s.staff_repo.IsStaff(request.GetStaffId())
	if err != nil {
		return nil, err
	}

	res := &v1.IsStaffResponse{
		IsStaff: isStaff,
	}

	return res, nil
}

func (s *StaffService) AddEventToUser(ctx context.Context, request *v1.AddEventToUserRequest) (*v1.AddEventToUserResponse, error) {
	isStaff, err := s.staff_repo.IsStaff(request.GetStaffUserId())
	if err != nil {
		return nil, err
	}

	if !isStaff {
		return nil, errors.New("Permission denied. Only staff user can perform this action.")
	}

	isTaken, err := s.user_repo.IsEventTaken(request.GetUserId(), request.GetEventId())
	if err != nil {
		return nil, err
	}

	if isTaken {
		return nil, errors.New("User has already been taken the event")
	}

	userEvent := event_ent.UserEvent{
		UserId:  request.GetUserId(),
		EventId: request.GetEventId(),
		TakenAt: time.Now().Unix(),
	}

	err = s.user_repo.AddEvent(userEvent)
	if err != nil {
		return nil, err
	}

	res := &v1.AddEventToUserResponse{
		Success: true,
	}

	return res, nil
}

func (s *StaffService) GenerateSignInToken(ctx context.Context, request *v1.GenerateSignInTokenRequest) (*v1.GenerateSignInTokenResponse, error) {
	isStaff, err := s.staff_repo.IsStaff(request.GetStaffUserId())
	if err != nil {
		return nil, err
	}

	if !isStaff {
		return nil, errors.New("Permission denied. Only staff user can perform this action.")
	}

	token := utils.GenToken(request.EventId + request.StaffUserId)
	tokenEntity := token_ent.Token{
		Id:      token,
		EventId: request.EventId,
		EndAt:   time.Now().Add(time.Minute * 15).Unix(), // token valids for 15 minutes
	}

	err = s.staff_repo.CreateToken(tokenEntity)
	if err != nil {
		return nil, err
	}

	res := &v1.GenerateSignInTokenResponse{
		Token: token,
	}

	return res, nil
}
