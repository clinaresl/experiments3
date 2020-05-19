// Este modulo traduce cadenas de texto al complicado lenguaje Kirgon
package mod2

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Módulo 2 activado ...")
	rand.Seed(time.Now().Unix())
}

// traduce cualquier cadena a Kirgon, lo creas o no
//  Translate("¡Hola!")
// devuelve el famoso saludo Kirgon H¡o!al
func Translate(cadena string) string {

	runes := []rune(cadena)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
