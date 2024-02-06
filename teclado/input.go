package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el numero 1: ")

	//Verificamos si el usuario ingreso algo por teclado
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Ingrese el numero 2: ")

	//Verificamos si el usuario ingreso algo por teclado
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
	}

	fmt.Println("Ingrese un mensaje: ")

	//Verificamos si el usuario ingreso algo por teclado
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
