package interfaces

//Se define lo que sabe hacer pero no como lo hace
type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
}
