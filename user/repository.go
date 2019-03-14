package user

// Data access layer. Encapsulates database access.
// Does not expose persistence related implementation details, e.g. SQL, to callers

type selectUserByID func(id string) string

func newSelectUserByID() selectUserByID {
	return func(id string) string {
		return "jan"
	}
}
