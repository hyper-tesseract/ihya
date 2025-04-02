package entity

import "strings"

type Persona struct {
	UserId      string
	PersonaName string
	title       string
	ThemeId     string
	IhyaWrapper bool
}

func (p *Persona) Title(name string) string {
	return strings.Join([]string{name, p.title}, " | ")
}

func NewFakePersona() Persona {
	return Persona{
		UserId:      "1",
		PersonaName: "hmuhammadazeem",
		title:       "Software Engineer",
		ThemeId:     "1",
		IhyaWrapper: true,
	}
}
