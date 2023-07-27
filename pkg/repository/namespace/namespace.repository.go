package namespace

import (
	namespace_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/namespace"
	namespace_repo "github.com/isd-sgcu/rpkm66-checkin/internal/repository/namespace"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllNamespace(result *[]*namespace_ent.Namespace) error
}

func NewRepository(db *gorm.DB) Repository {
	return namespace_repo.NewRepository(db)
}
