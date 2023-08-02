package staff

import (
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	staff_repo "github.com/isd-sgcu/rpkm66-checkin/internal/repository/staff"
	"gorm.io/gorm"
)

type Repository interface {
	IsStaff(userId string) error
	CreateToken(token token_ent.Token) error
}

func NewRepository(db *gorm.DB) Repository {
	return staff_repo.NewRepository(db)
}
