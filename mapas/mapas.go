package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Argentina"])

	//Mapas con llave int, ejemplo por ID
	usuarios := map[int]string{
		4445: "Emilio",
		7458: "Victoria",
	}

	fmt.Println(usuarios[4445])

	delete(usuarios, 4445)
	fmt.Println(usuarios[4445])
}
