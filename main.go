package main

import (
	ejerinterfaces "src/github.com/faridEmilio/go-proyects/ejer_interfaces"
	"src/github.com/faridEmilio/go-proyects/modelos"
)

func main() {
	/*variables.MuestroEnteros()
	estado, texto := variables.ConviertoATexto(15)
	fmt.Println(estado)
	fmt.Println(texto)

	numero, mensaje := ejercicios.Convertidor("400")
	fmt.Println(numero)
	fmt.Println(mensaje)


	teclado.IngresoNumeros()

	fmt.Println("Calculadora")
	ejercicios.Calculadora()
	*/
	//files.GrabaTabla()

	//files.ConcatenaTabla()

	//files.LeoArchivo()

	//funciones.Calculos()

	//funciones.LlamarClosure()

	//funciones.Exponencia(4)

	//arreglos_slices.MuestroArreglos()

	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()

	usuario := new(modelos.Hombre)
	maria := new(modelos.Mujer)

	ejerinterfaces.HumanosRespirando(usuario)

	ejerinterfaces.HumanosRespirando(maria)

}
