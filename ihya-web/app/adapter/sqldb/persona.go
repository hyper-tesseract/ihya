package sqldb

import (
	"fmt"
	entity "ihya-web/app/entity"
	userland "ihya-web/app/usecase/userland"
)

var _ userland.PersonaRepo = (*FakePersonaRepo)(nil)

type FakePersonaRepo struct {
	fakedb map[string]entity.Persona
}

func (f *FakePersonaRepo) GetPersona(name string) (entity.Persona, error) {
	persona, exists := f.fakedb[name]
	if !exists {
		return entity.Persona{}, fmt.Errorf("%s not found", name)
	}
	return persona, nil
}

func NewFakePersonaRepo() FakePersonaRepo {
	data := map[string]entity.Persona{}
	data["hmuhammadazeem"] = entity.NewFakePersona()
	return FakePersonaRepo{fakedb: data}
}
