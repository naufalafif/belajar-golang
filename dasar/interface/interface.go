package main

import "fmt"

type bot interface {
	getSalam() string
}

type botInggris struct{}

type botSpayol struct{}

// tidak perlu deklarasi objeknya karena tidak digunakan func (i botInggris) ...
func (botInggris) getSalam() string {
	return "Halo"
}

func (botSpayol) getSalam() string {
	return "Hola"
}

func main() {
	objekInggris := botInggris{}
	objekSpanyol := botSpayol{}

	printSalam(objekInggris)
	printSalam(objekSpanyol)
}

func printSalam(b bot) {
	fmt.Println(b.getSalam())
}
