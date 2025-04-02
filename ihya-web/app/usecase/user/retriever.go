package user

import "ihya-web/app/entity"

type UserRetriever struct {
	userRepo UserRetrieverRepo
}

func (u *UserRetriever) RetrieveUser(name string) (entity.User, error) {
	user, err := u.userRepo.GetUserById(name)
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRetriever(repo UserRetrieverRepo) UserRetriever {
	return UserRetriever{userRepo: repo}
}
