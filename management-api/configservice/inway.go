//nolint:dupl // service and inway structs look the same
package configservice

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/configapi"
)

// CreateInway creates a new inway
func (s *ConfigService) CreateInway(ctx context.Context, inway *configapi.Inway) (*configapi.Inway, error) {
	logger := s.logger.With(zap.String("name", inway.Name))
	logger.Info("rpc request CreateInway")

	err := inway.Validate()
	if err != nil {
		logger.Error("invalid inway", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid inway: %s", err))
	}

	err = s.configDatabase.CreateInway(ctx, inway)
	if err != nil {
		logger.Error("error creating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return inway, nil
}

// GetInway returns a specific inway
func (s *ConfigService) GetInway(ctx context.Context, req *configapi.GetInwayRequest) (*configapi.Inway, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request GetInway")

	inway, err := s.configDatabase.GetInway(ctx, req.Name)
	if err != nil {
		logger.Error("error getting inway from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if inway == nil {
		logger.Warn("inway not found")
		return nil, status.Error(codes.NotFound, "inway not found")
	}

	return inway, nil
}

// UpdateInway updates an existing inway
func (s *ConfigService) UpdateInway(ctx context.Context, req *configapi.UpdateInwayRequest) (*configapi.Inway, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request UpdateInway")

	err := req.Inway.Validate()
	if err != nil {
		logger.Error("invalid inway", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid inway: %s", err))
	}

	err = s.configDatabase.UpdateInway(ctx, req.Name, req.Inway)
	if err != nil {
		logger.Error("error updating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req.Inway, nil
}

// DeleteInway deletes a specific inway
func (s *ConfigService) DeleteInway(ctx context.Context, req *configapi.DeleteInwayRequest) (*configapi.Empty, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteInway")

	err := s.configDatabase.DeleteInway(ctx, req.Name)

	if err != nil {
		logger.Error("error deleting inway in DB", zap.Error(err))
		return &configapi.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &configapi.Empty{}, nil
}

// ListInways returns a list of inways
func (s *ConfigService) ListInways(ctx context.Context, req *configapi.ListInwaysRequest) (*configapi.ListInwaysResponse, error) {
	s.logger.Info("rpc request ListInways")

	inways, err := s.configDatabase.ListInways(ctx)
	if err != nil {
		s.logger.Error("error getting inway list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	services, err := s.configDatabase.ListServices(ctx)
	if err != nil {
		s.logger.Error("error getting services  from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if len(services) > 0 {
		for _, i := range inways {
			for _, service := range services {
				if contains(service.Inways, i.Name) {
					i.Services = append(i.Services, &configapi.Inway_Service{Name: service.Name})
				}
			}
		}
	}

	return &configapi.ListInwaysResponse{
		Inways: inways,
	}, nil
}