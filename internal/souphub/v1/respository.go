package souphubv1

import (
	"context"

	"github.com/rosemound/souphub/internal/domain/models"
)

type Repository struct {
	servers models.GameServers
}

func NewRepository(srvs models.GameServers) (*Repository, error) {
	return &Repository{servers: srvs}, nil
}

func (r *Repository) Create(ctx context.Context, addr string, srv *models.GameServer) error {
	if !r.IsExists(ctx, addr) {
		r.servers[addr] = srv
	}

	return nil
}

func (r *Repository) FindAll(ctx context.Context) models.GameServers {
	return r.servers
}

func (r *Repository) CreateAll(ctx context.Context, srvs models.GameServers) error {
	for addr, meta := range srvs {
		if r.IsExists(ctx, addr) {
			continue
		}

		r.servers[addr] = meta
	}

	return nil
}

func (r *Repository) IsExists(ctx context.Context, addr string) bool {
	_, ok := r.servers[addr]
	return ok
}