package entity

type PageData struct {
	Bio         AuthorBio
	ThemeRoot   string
	PageSlug    string
	SiteTitle   string
	IhyaWrapper bool
}

type AuthorBio struct {
	Name     string
	Title    string
	UserName string
	HeadLine string
}
