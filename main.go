package main

import (
	"github.com/steinfletcher/func-dependency-injection-go/user"
	"log"
	"net/http"
)

// run the following then call `curl -v localhost:8080/user`

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/user", user.Get())
	log.Fatal(http.ListenAndServe(":8080", router))
}
