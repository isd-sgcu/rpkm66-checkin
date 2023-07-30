package token

import (
	token_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{
		db,
	}
}

func (r *TokenRepository) GetTokenInfo(token string, result *token_ent.Token) error {
	return r.db.Model(&token_ent.Token{}).First(result, "id = ?", token).Error
}
