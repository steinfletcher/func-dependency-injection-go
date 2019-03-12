package user

import "net/http"

// Compile time dependency injection. Simple wiring without any frameworks or magic

func Get() http.HandlerFunc {
	repository := newSelectUserByID()
	service := newGetUser(repository)
	return newGetUserHandler(service)
}
