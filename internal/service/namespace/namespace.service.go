package namespace

import (
	"context"

	namespace_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/namespace"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/namespace/v1"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/namespace"
)

var _ v1.NamespaceServiceServer = &serviceImpl{}

type serviceImpl struct {
	v1.UnimplementedNamespaceServiceServer
	repo namespace.Repository
}

func NewService(repo namespace.Repository) v1.NamespaceServiceServer {
	return &serviceImpl{
		v1.UnimplementedNamespaceServiceServer{},
		repo,
	}
}

func (s *serviceImpl) GetAllNamespaces(ctx context.Context, request *v1.GetAllNamespacesRequest) (*v1.GetAllNamespacesResponse, error) {
	var namespaces []*namespace_ent.Namespace

	err := s.repo.GetAllNamespace(&namespaces)
	if err != nil {
		return nil, err
	}

	arr := make([]*v1.Namespace, len(namespaces))
	for _, namespace := range namespaces {
		arr = append(arr, namespace.ToProto())
	}

	response := &v1.GetAllNamespacesResponse{
		Namespaces: arr,
	}

	return response, nil
}
