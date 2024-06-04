package methods

import "strings"

// The code defines a struct type User with a single field named name of type string.
// @property {string} name - The `User` struct has a single property `name` of type `string`.
type User struct {
	Name string
}

// The `Greet()` method is a method defined on the `User` struct in Go. It returns a string that
// concatenates the "Hello" string with the `name` field of the `User` struct instance on which the
// method is called.
func (user User) Greet() string {
	return "Hello " + user.Name
}

// The `Uppername()` method is defined on the `User` struct in Go. It converts the `name` field of the
// `User` struct instance to uppercase using the `strings.ToUpper()` function. However, there is a
// mistake in the method signature. The method should return a string since it is modifying the `name`
// field. The correct method signature should be:
func (user User) Uppername() {
	user.Name = strings.ToUpper(user.Name)
}

// The `SetName` method in the code snippet is defined on the `User` struct in Go. It takes a `name`
// parameter of type string and sets the `name` field of the `User` struct instance to the value passed
// as the parameter. This method allows you to update or change the value of the `name` field for a
// specific `User` struct instance.
func (user *User) SetName(name string) {
	user.Name = name
}
