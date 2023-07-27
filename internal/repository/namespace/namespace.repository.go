package namespace

import (
	namespace_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/namespace"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{
		db,
	}
}

func (r *repositoryImpl) GetAllNamespace(result *[]*namespace_ent.Namespace) error {
	return r.db.Model(&namespace_ent.Namespace{}).Find(result).Error
}
