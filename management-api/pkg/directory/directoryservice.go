package directory

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/pkg/environment"
)

type Service struct {
	logger      *zap.Logger
	environment *environment.Environment

	directoryClient *Client
}

var inwayStateToDirectoryStatus = map[InwayState]DirectoryService_Status{
	InwayStateUnknown: DirectoryService_unknown,
	InwayStateUp:      DirectoryService_up,
	InwayStateDown:    DirectoryService_down,
}

func NewDirectoryService(logger *zap.Logger, e *environment.Environment, directoryClient *Client) *Service {
	return &Service{
		logger:          logger,
		environment:     e,
		directoryClient: directoryClient,
	}
}

// ListServices returns all services known to the directory except those with the same organization
func (s Service) ListServices(context.Context, *Empty) (*ListServicesResponse, error) {
	s.logger.Info("rpc request ListServices")

	listServices, err := s.directoryClient.ListServices()
	if err != nil {
		s.logger.Error("error getting services list from directory", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "directory error")
	}

	services := []*DirectoryService{}

	for _, service := range listServices {
		if service.OrganizationName == s.environment.OrganizationName {
			continue
		}

		services = append(services, convertDirectoryService(service))
	}

	return &ListServicesResponse{Services: services}, nil
}

// GetOrganizationService returns a specific service of and organization
func (s Service) GetOrganizationService(_ context.Context, request *GetOrganizationServiceRequest) (*DirectoryService, error) {
	logger := s.logger.With(zap.String("organizationName", request.OrganizationName), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request GetOrganizationService")

	service, err := s.getService(logger, request.OrganizationName, request.ServiceName)
	if err != nil {
		return nil, err
	}

	return convertDirectoryService(service), nil
}

func (s Service) getService(logger *zap.Logger, organizationName, serviceName string) (*InspectionAPIService, error) {
	services, err := s.directoryClient.ListServices()
	if err != nil {
		return nil, status.Error(codes.Internal, "directory not available")
	}

	for _, s := range services {
		if s.OrganizationName == organizationName && s.Name == serviceName {
			return s, nil
		}
	}

	logger.Warn("service not found")

	return nil, status.Error(codes.NotFound, "service not found")
}

// RequestAccessToService records an access request and sends it to the organization
func (s Service) RequestAccessToService(ctx context.Context, request *RequestAccessToServiceRequest) (*Empty, error) {
	logger := s.logger.With(zap.String("organizationName", request.OrganizationName), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request RequestAccessToService")

	service, err := s.getService(logger, request.OrganizationName, request.ServiceName)
	if err != nil {
		return nil, err
	}

	logger = logger.With(zap.Any("service", service))

	logger.Debug("Check if permission is required, fail otherwise")

	logger.Debug("persist access request")

	logger.Debug("send access request to inway")

	return &Empty{}, nil
}

func DetermineDirectoryServiceStatus(inways []*Inway) DirectoryService_Status {
	serviceStatus := DirectoryService_unknown

	if len(inways) == 0 {
		return serviceStatus
	}

	stateMap := map[InwayState]int{}

	for _, i := range inways {
		stateMap[i.State]++
	}

	if len(stateMap) > 1 {
		return DirectoryService_degraded
	}

	for state := range stateMap {
		serviceStatus = inwayStateToDirectoryStatus[state]
	}

	return serviceStatus
}

func convertDirectoryService(s *InspectionAPIService) *DirectoryService {
	serviceStatus := DetermineDirectoryServiceStatus(s.Inways)

	return &DirectoryService{
		ServiceName:          s.Name,
		OrganizationName:     s.OrganizationName,
		APISpecificationType: s.APISpecificationType,
		Status:               serviceStatus,
	}
}