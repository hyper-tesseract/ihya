package sqldb

import (
	"fmt"
	entity "ihya-web/app/entity"
	ssr "ihya-web/app/usecase/ssr"
)

var _ ssr.ThemeRepo = (*FakeThemeRepo)(nil)

type FakeThemeRepo struct {
	fakedb map[string]entity.Theme
}

func (f *FakeThemeRepo) GetTheme(themeId string) (entity.Theme, error) {
	theme, exists := f.fakedb[themeId]
	if !exists {
		return entity.Theme{}, fmt.Errorf("Theme %s not found", themeId)
	}
	return theme, nil
}

func NewFakeThemeRepo() FakeThemeRepo {
	data := map[string]entity.Theme{}
	data["1"] = entity.Theme{
		ThemeId:       "1",
		Root:          "porto",
		PageTemplates: []string{"porto-page-portfolio"},
	}
	return FakeThemeRepo{fakedb: data}
}
