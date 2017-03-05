// This sample code implement a simple web service.
package main

import (
	"log"
	"net/http"

	"github.com/dandawong/golang/test/handlers"
)

// main
func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}