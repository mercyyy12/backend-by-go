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

func (e email) send(msg string) string {
	return fmt.Sprintf("Email sent to %s: %s", e.email, msg)
}

func (s sms) send(msg string) string {
	return fmt.Sprintf("SMS sent to %s: %s", s.number, msg)
}

func notify(n notifier, msg string) {
	fmt.Println(n.send(msg))
}

func main() {
	em := email{"mercyyy@gmail.com"}
	notify(em, "santosh")

	sm := sms{"98545"}
	notify(sm, "santosh89654")
}
