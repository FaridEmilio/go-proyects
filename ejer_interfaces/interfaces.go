package ejerinterfaces

import (
	"fmt"
	"src/github.com/faridEmilio/go-proyects/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {
	humano.Comer()
	humano.Respirar()
	fmt.Printf("Soy un/a %s, y estoy comiendo \n", humano.Sexo())
}
