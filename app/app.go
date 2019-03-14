package app

import (
	"github.com/steinfletcher/func-dependency-injection-go/user"
	"net/http"
)

type App struct {
	Router *http.ServeMux
}

func New() *App {
	router := http.NewServeMux()
	router.HandleFunc("/user", user.Get())
	return &App{Router: router}
}
