package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rosemound/souphub/configs"
	dtosv1 "github.com/rosemound/souphub/internal/domain/dtos/v1"
	souphubv1 "github.com/rosemound/souphub/internal/souphub/v1"
)

func main() {
	var err error
	var config configs.Config

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover: %v", r)
		}
	}()

	if err = configs.Get(&config); err != nil {
		panic(err)
	}

	mux := newRouter(&config)

	if mux, err = souphubv1.RegisterSoupHubRouter(mux, souphubv1.RouterConfig{
		Service: souphubv1.ServiceConfig{Hub: dtosv1.Hub{
			Name:        config.Name,
			Description: config.Description,
			Company:     config.Company,
			Servers:     config.Servers,
		}},
	}); err != nil {
		panic(err)
	}

	s := http.Server{
		Addr:    ":" + config.HttpPort,
		Handler: mux,
	}

	if err = s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func newRouter(config *configs.Config) *chi.Mux {
	router := chi.NewMux()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "User-Agent"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(AccessTokenAuth(config.AccessToken))

	if config.IsDebug() {
		router.Use(middleware.Logger)
	}

	return router
}

func AccessTokenAuth(accessToken string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()

			if access := q.Get("access_token"); access != accessToken {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
