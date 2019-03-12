package user

import (
	"encoding/json"
	"net/http"
)

// Protocol layer. Reads the http request and writes the response
// Deals with http, cookies, headers, request validation etc.
// Does not perform business logic or call out to external services

func newGetUserHandler(getUser getUser) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		u := getUser("1234")
		data, _ := json.Marshal(u)

		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(data)
	}
}
