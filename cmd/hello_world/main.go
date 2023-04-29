package main

import (
	"fmt"
	"hello_world_go_webapp/internal/app"
	"hello_world_go_webapp/internal/middleware"
	"net/http"
)

func main() {
	port := 8080
	mux := app.SetupRoutes()
	loggedMux := middleware.LoggingMiddleware(mux)
	http.ListenAndServe(fmt.Sprintf(":%v", port), loggedMux)
}
