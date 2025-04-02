package entity

import "strings"

type User struct {
	UserId       string
	firstName    string
	lastName     string
	Title        string
	Bio          string
	ProfileImage string
}

func (u *User) Name() string {
	return strings.Join([]string{u.firstName, u.lastName}, " ")
}

func NewFakeUser() User {
	return User{
		UserId:       "1",
		firstName:    "Muhammad",
		lastName:     "Azeem",
		Title:        "Software Engineer",
		Bio:          "Software Engineer with passion for Distributed Systems. I enjoy creating tools that make life easier for people.",
		ProfileImage: "",
	}
}
