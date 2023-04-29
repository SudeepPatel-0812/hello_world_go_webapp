package app

import (
	"hello_world_go_webapp/internal/routes"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	return mux
}
