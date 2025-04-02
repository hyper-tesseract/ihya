package ssr

import "ihya-web/app/entity"

func PreparePersonaTemplateMapping(
	persona entity.Persona,
	theme entity.Theme,
	user entity.User,
) entity.PageData {
	bio := entity.AuthorBio{Name: user.Name(), Title: user.Title, HeadLine: user.Bio}

	return entity.PageData{
		Bio:         bio,
		ThemeRoot:   theme.Root,
		PageSlug:    theme.PageTemplates[0],
		SiteTitle:   persona.Title(user.Name()),
		IhyaWrapper: persona.IhyaWrapper,
	}
}
