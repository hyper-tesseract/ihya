package sqldb

import (
	"fmt"
	entity "ihya-web/app/entity"
	"ihya-web/app/usecase/user"
)

var _ user.UserRetrieverRepo = (*FakeUserRetrieverRepo)(nil)

type FakeUserRetrieverRepo struct {
	fakedb map[string]entity.User
}

func (f *FakeUserRetrieverRepo) GetUserById(userId string) (entity.User, error) {
	user, exists := f.fakedb[userId]
	if !exists {
		return entity.User{}, fmt.Errorf("%s not found", userId)
	}
	return user, nil
}

func NewFakeUserRetrieverRepo() FakeUserRetrieverRepo {
	data := map[string]entity.User{}
	data["1"] = entity.NewFakeUser()
	return FakeUserRetrieverRepo{fakedb: data}
}
