package souphubv1

import (
	"context"
	"sync"
	"time"

	"github.com/rosemound/souphub/internal/domain/models"
)

type Registry struct {
	masters models.Masters

	mut sync.RWMutex
}

func NewRegistry() *Registry {
	reg := &Registry{models.Masters{}, sync.RWMutex{}}
	go reg.expiration()
	return reg
}

func (r *Registry) Create(ctx context.Context, addr models.MasterToken, meta *models.Master) error {
	meta.Expiration = time.Now().Add(time.Duration(time.Hour * 48)).Unix()

	r.mut.Lock()
	defer r.mut.Unlock()

	r.masters[addr] = meta

	return nil
}

func (r *Registry) CreateAll(ctx context.Context, masters models.Masters) error {
	for addr, meta := range masters {
		if r.IsExists(ctx, addr) {
			continue
		}

		r.Create(ctx, addr, meta)
	}

	return nil
}

func (r *Registry) Get(ctx context.Context, token models.MasterToken) *models.Master {
	r.mut.RLock()
	defer r.mut.RUnlock()
	return r.masters[token]
}

func (r *Registry) GetAll(ctx context.Context) models.Masters {
	r.mut.RLock()
	defer r.mut.RUnlock()
	return r.masters
}

func (r *Registry) Delete(ctx context.Context, token models.MasterToken) {
	r.mut.Lock()
	defer r.mut.Unlock()
	delete(r.masters, token)
}

func (r *Registry) IsExists(ctx context.Context, token models.MasterToken) bool {
	_, ok := r.masters[token]
	return ok
}

func (r *Registry) expiration() {
	ctx := context.Background()

	for {
		for addr, reg := range r.GetAll(ctx) {
			if reg == nil || reg.Expiration <= time.Now().Unix() {
				r.Delete(ctx, addr)
			}
		}

		time.Sleep(time.Duration(time.Minute))
	}
}
