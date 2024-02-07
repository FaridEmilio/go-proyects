package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Segundo mensaje")
}

// Aborta el programa con un mensaje por consola
func EjemploPanic() {

	defer func() {
		reco := recover()

		if reco != nil {
			log.Fatalf("Ocurrio un error que generó Panic \n %v", reco)
		}
	}()

	a := 1

	if a == 1 {
		panic("Se encontro un error")
	}
}
