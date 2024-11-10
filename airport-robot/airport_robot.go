package airportrobot

import "fmt"

type Greeter interface {
	Language() string
	Greet(name string) string
}

// SayHello greets the given name in the language of the greeter
func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", greeter.Language(), greeter.Greet(name))
}

type Italian struct{}

// Language returns the language of the greeter
func (i Italian) Language() string {
	return "Italian"
}

// Greet greets the given name
func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s", name)
}

type Portuguese struct{}

// Language returns the language of the greeter
func (p Portuguese) Language() string {
	return "Portuguese"
}

// Greet greets the given name
func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s", name)
}
