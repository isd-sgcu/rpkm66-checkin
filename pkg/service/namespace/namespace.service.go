package namespace

import (
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/namespace/v1"
	namespace_svc "github.com/isd-sgcu/rpkm66-checkin/internal/service/namespace"
	namespace_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/namespace"
)

func NewService(repo namespace_repo.Repository) v1.NamespaceServiceServer {
	return namespace_svc.NewService(repo)
}
