package variables

import (
	"strconv"
)

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)

	return true, texto
}
