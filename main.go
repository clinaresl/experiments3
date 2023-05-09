// Ejercicio de inicialización de paquetes
package main

import (
	"fmt"

	"github.com/clinaresl/experiments3/mod2"
)

func init() {
	fmt.Println("Ready to launch ...")
}

func main() {
	saludo := "Liftoff!"
	fmt.Println(saludo)
	fmt.Println(mod2.Translate("Arecibo message sent"))
}
