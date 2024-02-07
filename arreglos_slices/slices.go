package arreglos_slices

import "fmt"

var slice []int
var arreglo [10]int = [10]int{6, 58, 45, 741, 58, 1, 25, 3, 69, 7}

func MuestroSlices() {

	fmt.Println(arreglo)

	//Un slice de ese vector desde la posicion 3 hasta el final
	porcion := arreglo[3:]

	fmt.Println(porcion)
}

func Capacidad() {
	// 5 elementos, 20 de capacidad (reserva en memoria esa capacidad)
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}
