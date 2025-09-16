package souphubv1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	dtosv1 "github.com/rosemound/souphub/internal/domain/dtos/v1"
	"github.com/go-chi/render"
)

type RouterConfig struct {
	Service ServiceConfig
}

type Router struct {
	chi.Router

	service *Service
}

func RegisterSoupHubRouter(c *chi.Mux, config RouterConfig) (*chi.Mux, error) {
	router, err := NewRounter(config)

	if err != nil {
		return nil, err
	}

	c.Mount("/souph", router)

	return c, nil
}

func NewRounter(config RouterConfig) (http.Handler, error) {
	service, err := NewService(config.Service)
	if err != nil {
		return nil, err
	}

	router := &Router{chi.NewRouter(), service}
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Post("/connect", router.Connect)
	router.Get("/share", router.Share)

	return router, nil
}

func (r *Router) Connect(w http.ResponseWriter, req *http.Request) {
	var data dtosv1.MasterHubConnect

	if err := render.Bind(req, &data); err != nil {
		render.Render(w, req, dtosv1.NewHttpErr(http.StatusBadRequest, err))
		return
	}

	body, err := r.service.Connect(req.Context(), &data)

	if err != nil {
		render.Render(w, req, dtosv1.NewHttpErr(http.StatusBadRequest, err))
		return
	}

	render.Status(req, http.StatusCreated)
	render.Render(w, req, body)
}

func (r *Router) Share(w http.ResponseWriter, req *http.Request) {
	var data dtosv1.Share

	if err := render.Bind(req, &data); err != nil {
		render.Render(w, req, dtosv1.NewHttpErr(http.StatusBadRequest, err))
		return
	}

	body, err := r.service.Share(req.Context(), &data)

	if err != nil {
		render.Render(w, req, dtosv1.NewHttpErr(http.StatusBadRequest, err))
		return
	}

	render.Status(req, http.StatusCreated)
	render.Render(w, req, body)
}
