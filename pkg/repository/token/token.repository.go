package token

import (
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	token_repo "github.com/isd-sgcu/rpkm66-checkin/internal/repository/token"
	"gorm.io/gorm"
)

type Repository interface {
	GetTokenInfo(token string, result *token_ent.Token) error
}

func NewRepository(db *gorm.DB) Repository {
	return token_repo.NewRepository(db)
}
