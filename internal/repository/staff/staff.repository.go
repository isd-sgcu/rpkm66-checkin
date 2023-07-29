package staff

import (
	event_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
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

func (r *StaffRepository) IsStaff(userId string, result *bool) error {
	var staff staff_ent.Staff
	err := r.db.Model(&staff_ent.Staff{}).Find(&staff, "user_id = ?", userId).Error
	if err != nil {
		return err
	}

	*result = staff.UserId != ""
	return err
}

func (r *StaffRepository) AddEventToUser(user_id string, event_id string) error {
	userEvent := event_ent.UserEvent{
		UserId:  user_id,
		EventId: event_id,
	}

	return r.db.Create(userEvent).Error
}

func (r *StaffRepository) CreateToken(token token_ent.Token) error {
	return r.db.Create(token).Error
}
