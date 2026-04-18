package goguide

import "fmt"

// PanicExample dispara um panic, que interrompe a execução normal.
func PanicExample() {
	panic("erro fatal")
}

// RecoverExample mostra como recuperar um panic com recover.
func RecoverExample() (recovered bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperado de panic:", r)
			recovered = true
		}
	}()
	PanicExample()
	return false
}
