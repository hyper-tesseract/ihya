package ssr

import "ihya-web/app/entity"

type ThemeRepo interface {
	GetTheme(themeId string) (entity.Theme, error)
}
