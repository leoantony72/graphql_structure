package repositories

import (
	"test/internal/model"

	"github.com/google/uuid"
)

func (t TestRepo) GetData(id string) *model.User {
	userData := model.User{}
	t.db.Where("id = ?", id).First(&userData)
	return &userData
}
func (t TestRepo) GetDataall() []*model.User {
	var users []*model.User
	t.db.Table("users").Select("*").Scan(&users)
	return users
}
func (t TestRepo) CreateUser(user *model.User) *model.User {
	id := uuid.New()
	user.Id = id.String()
	t.db.Create(&user)
	return user
}
