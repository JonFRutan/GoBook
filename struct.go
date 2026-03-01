package main

import (
	"fmt"
	"math/rand"
)

type user struct {
	Name string      //public field name
	Id int			 //When a field is capitalized, it's public
	pin int          //private field pin, indicated by the fact that it is lowercase
}

//A struct method, requires you to call it from a pointer reference to a user object
func (u *user) updatePin(pin int) {
	//fmt.Printf("Updating pin: %d\n", pin)
	u.pin = pin
}

func newUser(name string) *user {
	u := user{Name: name, Id: rand.Intn(100)}
	return &u
}

func main() {
	//an array is a statically sized set of elements.
	userArray := [7]string{"Joe", "Bobby", "Robotnik", "Mario", "Leopold", "Hulk", "Pringles"}
	//a slice is a dynamically sized set of elements.
	//In this instance, we make it an array of pointers to user structs.
	var users []*user

	//we use ranges to iterate over the length of an array
	for i := range userArray {
		newUser := newUser(userArray[i])
		//append adds new entires to a slice
		users = append(users, newUser)
	}
	//this is a proper "for-each" loop
	//the first value is our index, the second is our value from the range
	//we use an underscore as an "anonymous" index if we don't want to use it, and don't want to throw errors.
	for _, v := range users {
		v.updatePin(rand.Intn(89999) + 10000)
		fmt.Printf("Username: %s | ID: %d | Pin: %d\n", v.Name, v.Id, v.pin)
	}
}