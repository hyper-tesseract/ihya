package user

import "ihya-web/app/entity"

type UserRetrieverRepo interface {
	GetUserById(userId string) (entity.User, error)
}
