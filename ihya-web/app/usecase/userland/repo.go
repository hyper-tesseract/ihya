package userland

import (
	"ihya-web/app/entity"
)

// Persona repo accesses users' web page information from storage, such as database.
type PersonaRepo interface {
	GetPersona(name string) (entity.Persona, error)
}
