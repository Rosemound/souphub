package dtosv1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rosemound/souphub/internal/domain/models"
)

var _ render.Renderer = (*Masters)(nil)

type Masters struct {
	Masters []*models.Master `json:"masters"`
}

// Render implements render.Renderer.
func (m *Masters) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
