package modelos

import (
	"time"
)

// Aca creamos la entidad User
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// Aca le podemos dar un "comportamiento"
// Este es un "constructor"
func (this *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.Status = status
}
