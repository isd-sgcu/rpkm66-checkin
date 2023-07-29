package staff

import (
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/staff/v1"
	staff_svc "github.com/isd-sgcu/rpkm66-checkin/internal/service/staff"
	staff_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/staff"
	user_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/user"
)

func NewService(staff_repo staff_repo.Repository, user_repo user_repo.Repository) v1.StaffServiceServer {
	return staff_svc.NewService(staff_repo, user_repo)
}
