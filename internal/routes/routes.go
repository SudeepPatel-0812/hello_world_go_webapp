package routes

import (
	"hello_world_go_webapp/internal/controllers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/say-hello", controllers.HelloWorldController)
	mux.HandleFunc("/say-bye", controllers.ByeWorldController)
}
