package souphubv1

import (
	"context"
	"errors"

	dtosv1 "github.com/rosemound/souphub/internal/domain/dtos/v1"
	"github.com/rosemound/souphub/internal/domain/models"
)

type ServiceConfig struct {
	dtosv1.Hub
}

type Service struct {
	registry Registry

	name    string
	desc    string
	company *models.Company

	repository *Repository
}

func NewService(config ServiceConfig) (*Service, error) {
	repo, err := NewRepository(config.Servers)

	if err != nil {
		return nil, err
	}

	return &Service{
		registry:   *NewRegistry(),
		name:       config.Name,
		desc:       config.Description,
		company:    config.Company,
		repository: repo,
	}, nil
}

func (s *Service) Connect(ctx context.Context, data *dtosv1.MasterHubConnect) (*dtosv1.MasterHubConnected, error) {

	for t, m := range data.Masters {
		s.registry.Create(ctx, t, m)
	}

	return &dtosv1.MasterHubConnected{Success: true}, nil
}

func (s *Service) Share(ctx context.Context, data *dtosv1.Share) (*dtosv1.Hub, error) {
	if ok := s.registry.IsExists(ctx, data.Token); !ok {
		return nil, errors.New("ERR_INVALID_DATA")
	}

	m := s.registry.Get(ctx, data.Token)
	
	srvs := s.repository.FindAll(ctx)

	buf := models.GameServers{}

	for _, s := range m.Addrs {
		buf[s] = srvs[s]
	}

	return &dtosv1.Hub{
		Name:        s.name,
		Description: s.desc,
		Company:     s.company,
		Servers:     buf,
	}, nil
}
