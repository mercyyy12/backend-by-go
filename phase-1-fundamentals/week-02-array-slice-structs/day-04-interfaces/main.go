package main

import "fmt"

type notifier interface {
	send(msg string) string
}

type email struct {
	email string
}
type sms struct {
	number string
}

// Implement send() for email
func (e email) send(msg string) string {
	return fmt.Sprintf("\nEmail sent to %s: %s\n", e.email, msg)
}

// Implement send() for sms
func (s sms) send(msg string) string {
	return fmt.Sprintf("\nSMS sent to %s: %s\n", s.number, msg)
}

func notify(n notifier, msg string) {

	// type switch
	switch v := n.(type) {
	case email:
		fmt.Println("Handled as Email", v.send(msg))
	case sms:
		fmt.Println("Handled as SMS", v.send(msg))
	default:
		fmt.Println("Unknown notifier type")
	}

}

func main() {
	var em notifier = email{"mercyyy@gmail.com"}
	notify(em, "How are you?")

	sm := sms{"98545"}
	notify(sm, "Yo! What's up?")

	// empty interface (when data is unknown)
	var x any
	// or
	// var x interface{}
	x = 8
	fmt.Println(x)
	x = "Mercyyy"
	fmt.Println(x)

	// type assertion (similar to maps)
	eml, ok := em.(email)
	if ok {
		fmt.Println(eml, ok)
	} else {
		fmt.Println(ok)
	}

	// Nil in interface
	var ni interface{} = nil // Interface is nil
	fmt.Println(ni == nil)

	// var e *email = nil // Interface holds a typed nil
	// var isnil notifier = e
	// fmt.Println(isnil == nil)

}
