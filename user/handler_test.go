package user

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/user/1234", nil)
	res := httptest.NewRecorder()
	getUserMock := func(id string) User {
		return User{
			Name: "jan",
			ID:   "1234",
		}
	}

	newGetUserHandler(getUserMock)(res, req)

	if res.Code != http.StatusOK {
		t.Fail()
	}
	if res.Body.String() != `{"name":"jan","id":"1234"}` {
		t.Fail()
	}
}
