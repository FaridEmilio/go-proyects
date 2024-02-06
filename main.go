package main

import (
	"fmt"
	"src/github.com/faridEmilio/go-proyects/variables"
)

func main() {
	variables.MuestroEnteros()
	estado, texto := variables.ConviertoATexto(15)
	fmt.Println(estado)
	fmt.Println(texto)
}
