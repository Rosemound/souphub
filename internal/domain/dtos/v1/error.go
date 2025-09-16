package dtosv1

import (
	"net/http"

	"github.com/go-chi/render"
)

var _ render.Renderer = (*HttpError)(nil)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Render implements render.Renderer.
func (h *HttpError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, h.Code)
	return nil
}

func NewHttpErr(code int, err error) render.Renderer {
	return &HttpError{Code: code, Message: err.Error(), Err: err}
}
