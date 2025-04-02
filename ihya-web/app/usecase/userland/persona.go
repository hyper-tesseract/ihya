package userland

import "ihya-web/app/entity"

type PersonaProvider struct {
	personaRepo PersonaRepo
}

func (p *PersonaProvider) RetrievePersona(name string) (entity.Persona, error) {
	persona, err := p.personaRepo.GetPersona(name)
	if err != nil {
		return persona, err
	}
	return persona, nil
}

func NewPersonaRenderer(repo PersonaRepo) PersonaProvider {
	return PersonaProvider{personaRepo: repo}
}
