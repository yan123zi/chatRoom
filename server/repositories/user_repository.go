package repositories

import "chatRoom/server/config"

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (r userRepository) FindOneByCondition() {
	return config.Db.Take("")
}
