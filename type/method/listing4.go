// https://play.golang.org/p/SByz5Esulx
// Sample program to show how to use an interface in Go.
package main

import (
	"fmt"
)

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	// cannot use u (type user) as type notifier in argument to sendNotification:
	// user does not implement notifier (notify method has pointer receiver)
	// panic
	// sendNotification(u)
	sendNotification(&u)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
