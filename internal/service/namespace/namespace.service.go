package namespace

import (
	"context"

	namespace_ent "github.com/isd-sgcu/rpkm66-checkin/internal/entity/namespace"
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/namespace/v1"
	"github.com/isd-sgcu/rpkm66-checkin/pkg/repository/namespace"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		log.Error().Err(err).
			Str("service", "checkin").
			Str("module", "GetAllNamespaces").
			Msg("Error while getting all namespaces")

		return nil, status.Error(codes.Internal, "Internal server error")
	}

	arr := make([]*v1.Namespace, len(namespaces))
	for i, namespace := range namespaces {
		arr[i] = namespace.ToProto()
	}

	response := &v1.GetAllNamespacesResponse{
		Namespaces: arr,
	}

	return response, nil
}
