package staff

import (
	"context"
	"errors"

	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/staff/v1"
	"github.com/isd-sgcu/rpkm66-checkin/internal/utils"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/staff"
)

type StaffService struct {
	v1.UnimplementedStaffServiceServer
	repo staff.Repository
}

func NewService(repo staff.Repository) v1.StaffServiceServer {
	return &StaffService{
		v1.UnimplementedStaffServiceServer{},
		repo,
	}
}

func (s *StaffService) IsStaff(ctx context.Context, request *v1.IsStaffRequest) (*v1.IsStaffResponse, error) {
	var isStaff bool
	err := s.repo.IsStaff(request.GetStaffId(), &isStaff)
	if err != nil {
		return nil, err
	}

	res := &v1.IsStaffResponse{
		IsStaff: isStaff,
	}

	return res, nil
}

func (s *StaffService) AddEventToUser(ctx context.Context, request *v1.AddEventToUserRequest) (*v1.AddEventToUserResponse, error) {
	var isStaff bool
	err := s.repo.IsStaff(request.GetStaffUserId(), &isStaff)
	if err != nil {
		return nil, err
	}

	if !isStaff {
		return nil, errors.New("Permission denied. Only staff user can perform this action.")
	}

	err = s.repo.AddEventToUser(request.UserId, request.EventId)
	if err != nil {
		return nil, err
	}

	res := &v1.AddEventToUserResponse{
		Success: true,
	}

	return res, nil
}

func (s *StaffService) GenerateSignInToken(ctx context.Context, request *v1.GenerateSignInTokenRequest) (*v1.GenerateSignInTokenResponse, error) {
	var isStaff bool
	err := s.repo.IsStaff(request.GetStaffUserId(), &isStaff)
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
		EndAt:   request.EndAt,
	}

	err = s.repo.CreateToken(tokenEntity)
	if err != nil {
		return nil, err
	}

	res := &v1.GenerateSignInTokenResponse{
		Token: token,
	}

	return res, nil
}