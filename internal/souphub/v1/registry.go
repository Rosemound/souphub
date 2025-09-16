package souphubv1

import (
	"context"

	"github.com/rosemound/souphub/internal/domain/models"
)

type Registry struct {
	masters models.Masters
}

func NewRegistry() *Registry {
	return &Registry{models.Masters{}}
}

func (r *Registry) Create(ctx context.Context, token string, meta *models.Master) error {
	r.masters[token] = meta

	return nil
}

func (r *Registry) CreateAll(ctx context.Context, masters models.Masters) error {
	for addr, meta := range masters {
		if r.IsExists(ctx, addr) {
			continue
		}

		r.masters[addr] = meta
	}

	return nil
}

func (r *Registry) Get(ctx context.Context, token string) *models.Master {
	return r.masters[token]
}

func (r *Registry) Delete(ctx context.Context, token string) {
	delete(r.masters, token)
}

func (r *Registry) IsExists(ctx context.Context, token string) bool {
	_, ok := r.masters[token]
	return ok
}
