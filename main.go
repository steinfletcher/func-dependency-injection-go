package main

import (
	"github.com/steinfletcher/func-dependency-injection-go/app"
	"log"
	"net/http"
)

// run main then call `curl -v localhost:8080/user`

func main() {
	newApp := app.New()
	log.Fatal(http.ListenAndServe(":8000", newApp.Router))
}
