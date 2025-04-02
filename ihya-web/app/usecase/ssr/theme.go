package ssr

import "ihya-web/app/entity"

type ThemeProvider struct {
	themeRepo ThemeRepo
}

func (p *ThemeProvider) RetrieveTheme(themeId string) (entity.Theme, error) {
	theme, err := p.themeRepo.GetTheme(themeId)
	if err != nil {
		return theme, err
	}
	return theme, nil
}

func NewThemeProvider(repo ThemeRepo) ThemeProvider {
	return ThemeProvider{themeRepo: repo}
}
