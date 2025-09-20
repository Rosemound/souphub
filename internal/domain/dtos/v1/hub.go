package dtosv1

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rosemound/souphub/internal/domain/models"
)

var _ render.Binder = (*MasterHubConnect)(nil)
var _ render.Binder = (*Share)(nil)
var _ render.Renderer = (*MasterHubConnected)(nil)
var _ render.Renderer = (*Hub)(nil)

type MasterHubConnect struct {
	Masters models.Masters `json:"masters"`
}

func (mhcon *MasterHubConnect) Bind(r *http.Request) error {
	if mhcon.Masters == nil {
		return errors.New("ERR_INVALID_DATA")
	}

	return nil
}

type MasterHubConnected struct {
	Success   bool             `json:"success"`
	Connected []*models.Master `json:"connected"`
}

func (mhcon *MasterHubConnected) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Share struct {
	Token string `json:"master_token"`
}

func (sh *Share) Bind(r *http.Request) error {
	return nil
}

type Hub struct {
	Name        string             `json:"name"`
	Description string             `json:"descritption,omitempty"`
	Company     *models.Company    `json:"company,omitempty"`
	Servers     models.GameServers `json:"servers,omitempty"`
}

func (h *Hub) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
