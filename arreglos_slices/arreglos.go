package arreglos_slices

import "fmt"

//Vector
var tabla [10]int

func MuestroArreglos() {
	tabla[5] = 100
	tabla[6] = 100

	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
}
