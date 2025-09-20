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

	done := []*models.Master{}

	for t, m := range data.Masters {
		if m.Host == "" || m.Name == "" {
			continue
		}

		m.Expiration = 0

		if err := s.registry.Create(ctx, t, m); err == nil {
			done = append(done, m)
		}
	}

	return &dtosv1.MasterHubConnected{Success: true, Connected: done}, nil
}

func (s *Service) Share(ctx context.Context, data *dtosv1.Share) (*dtosv1.Hub, error) {
	mtoken := models.MasterToken(data.Token)

	if ok := s.registry.IsExists(ctx, mtoken); !ok {
		return nil, errors.New("ERR_INVALID_DATA")
	}

	m := s.registry.Get(ctx, mtoken)

	srvs := s.repository.FindAll(ctx)

	buf := models.GameServers{}

	for _, s := range m.Addrs {
		buf[models.GameServerAddr(s)] = srvs[models.GameServerAddr(s)]
	}

	return &dtosv1.Hub{
		Name:        s.name,
		Description: s.desc,
		Company:     s.company,
		Servers:     buf,
	}, nil
}

func (s *Service) Masters(ctx context.Context) (*dtosv1.Masters, error) {
	var out dtosv1.Masters

	for _, v := range s.registry.GetAll(ctx) {
		out.Masters = append(out.Masters, v)
	}

	return &out, nil
}
