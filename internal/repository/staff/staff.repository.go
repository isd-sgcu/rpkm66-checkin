package staff

import (
	staff_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/staff"
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	"gorm.io/gorm"
)

type StaffRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{
		db,
	}
}

func (r *StaffRepository) IsStaff(userId string) (bool, error) {
	var staff staff_ent.Staff
	err := r.db.Model(&staff_ent.Staff{}).First(&staff, "user_id = ?", userId).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *StaffRepository) CreateToken(token token_ent.Token) error {
	return r.db.Create(token).Error
}
