package namespace

import (
	namespace_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/namespace"
	"gorm.io/gorm"
)

type NamespaceRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *NamespaceRepository {
	return &NamespaceRepository{
		db,
	}
}

func (r *NamespaceRepository) GetAllNamespace(result *[]*namespace_ent.Namespace) error {
	return r.db.Model(&namespace_ent.Namespace{}).Find(result).Error
}
