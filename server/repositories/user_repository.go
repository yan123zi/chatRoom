package repositories

import (
	"chatRoom/server/config"
	"chatRoom/server/model"
)

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (r userRepository) Create(user *model.User) error {
	return config.Db.Create(user).Error
}

func (r userRepository) Take(condition ...interface{}) *model.User {
	resp := &model.User{}
	if err := config.Db.Take(resp, condition).Error; err != nil {
		return nil
	}
	return resp
}

func (r *userRepository) FindOneByUserName(username string) *model.User {
	return r.Take("username=?", username)
}
