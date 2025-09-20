package dtosv1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rosemound/souphub/internal/domain/models"
)

var _ render.Renderer = (*Master)(nil)

type Master models.Master

// Master implements render.Renderer.
func (m *Master) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

