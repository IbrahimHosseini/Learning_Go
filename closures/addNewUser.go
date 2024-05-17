package closures

type User struct {
	Firstname string
	Lastname  string
}

func NewUser(firstname, lastname string) func() User {
	user := User{Firstname: firstname, Lastname: lastname}

	return func() User {
		return user
	}
}
