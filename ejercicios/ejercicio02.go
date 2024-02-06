package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var mensaje string
var result string

func Calculadora() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero para conocer sus multiplos: ")

		//Verificamos si el usuario ingreso algo por teclado
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	var i int
	for i = 1; i <= 10; i++ {
		result += fmt.Sprintf(" %d x %d = %d \n", i, numero, i*numero)
	}

	return result
}
