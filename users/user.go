package users

import (
	"fmt"
	"src/github.com/faridEmilio/go-proyects/modelos"
	"time"
)

func AltaUsuario() {

	usuario := new(modelos.User)
	usuario.AddUser(458, "Emilio", time.Now(), true)
	fmt.Println(usuario)

}
