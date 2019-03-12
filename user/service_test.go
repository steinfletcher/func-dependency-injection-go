package user

import (
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {
	selectUserMock := func(id string) string {
		return "jan"
	}

	user := newGetUser(selectUserMock)("12345")

	if !reflect.DeepEqual(User{ID: "12345", Name: "jan"}, user) {
		t.Fail()
	}
}

