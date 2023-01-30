package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(lang string) string
}

type Italian struct {
}
type Portuguese struct {
}

func (l Italian) LanguageName() string {
	return "Italian"
}

func (l Italian) Greet(name string) string {
	return fmt.Sprintf("I can speak Italian: Ciao %s!", name)
}

func (l Portuguese) LanguageName() string {
	return "Portuguese"
}

func (l Portuguese) Greet(name string) string {
	return fmt.Sprintf("I can speak Portuguese: Olá %s!", name)
}

func SayHello(name string, g Greeter) string {
	return g.Greet(name)
}
