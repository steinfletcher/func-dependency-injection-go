package user

// Domain layer. Contains application business logic.
// Should not see any direct references to protocol layer, e.g. no http or SQL

type User struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type getUser func(id string) User

func newGetUser(selectUser selectUserByID) getUser {
	return func(id string) User {
		name := selectUser(id)
		return User{
			ID:   id,
			Name: name,
		}
	}
}
