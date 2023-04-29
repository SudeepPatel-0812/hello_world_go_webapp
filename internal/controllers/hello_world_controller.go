package controllers

import (
	"fmt"
	"net/http"
)

func HelloWorldController(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
