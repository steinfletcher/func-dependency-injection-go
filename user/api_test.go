package user_test

import (
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/func-dependency-injection-go/app"
	"testing"
)

// a black box test of the api using https://github.com/steinfletcher/apitest

func TestGetUserApi(t *testing.T) {
	apitest.New("gets the user").
		Handler(app.New().Router).
		Get("/user").
		Expect(t).
		Body(`{"name":"jan","id":"1234"}`).
		End()
}
