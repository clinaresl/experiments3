// Módulo de mi primer programa en Go, no solo saluda, sino que además lo
// traduce a Kirgon
package main

import (
	"fmt"

	"github.com/clinaresl/experiments3/mod2"
)

func init() {
	fmt.Println("Ready to launch ...")
}

func main() {
	saludo := "¡Hola!"
	fmt.Println(saludo)
	fmt.Println(mod2.Translate(saludo))
}
