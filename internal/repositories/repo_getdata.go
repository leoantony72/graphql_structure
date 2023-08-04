package repositories

import (
	"test/internal/model"
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
