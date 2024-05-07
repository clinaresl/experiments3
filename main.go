// Ejercicio de inicialización de paquetes
package main

import (
	"fmt"

	"github.com/clinaresl/experiments3/mod2"
	_ "github.com/clinaresl/experiments3/mod3"
	"github.com/clinaresl/experiments3/payload"
)

func init() {
	fmt.Println("Main: Ready to launch ...")
}

func main() {
	payload.Ready()
	saludo := "Main: Liftoff!"
	fmt.Println(saludo)
	fmt.Println(mod2.Translate("Arecibo message sent"))
}
