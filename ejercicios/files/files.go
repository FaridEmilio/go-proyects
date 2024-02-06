package files

import (
	"bufio"
	"fmt"
	"os"
	"src/github.com/faridEmilio/go-proyects/ejercicios"
)

var fileName string = "./ejercicios/files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Calculadora()

	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}

	fmt.Fprint(archivo, texto)
	archivo.Close()

}

func ConcatenaTabla() {
	var texto string = ejercicios.Calculadora()

	//Condicion = Si NO ocurrio el append, pasa esto
	if !Append(texto) {
		fmt.Println("Error al concatenar archivos")
	}
}

// Flags
func Append(texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false

	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(registro)
	}

	archivo.Close()
}
