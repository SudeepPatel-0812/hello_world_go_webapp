package controllers

import (
	"fmt"
	"net/http"
)

func ByeWorldController(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Bye World! yes this is working")
}
