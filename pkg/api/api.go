package api

import (
	pb "dynamic-equity/proto/dynamicequity"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

// New muxer
func New() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	authServer := pb.NewAuthServer(&AuthController{}, nil)
	pieServer := pb.NewPieServer(&PieController{}, nil)

	r.Mount(authServer.PathPrefix(), authServer)
	r.Mount(pieServer.PathPrefix(), pieServer)

	return r
}
