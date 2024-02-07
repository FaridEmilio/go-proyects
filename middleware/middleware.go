package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)

	result = operacionesMidd(restar)(10, 3)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

//Estoy agarrando la funcion, la puedo prosesar y retorno lo mismo
// como si no hubiese pasado nada
func operacionesMidd(f func(int, int) int) func(int, int) int {

	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
