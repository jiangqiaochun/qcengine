package servicepackage

import "ipadgrpc/src/model"

type UserService struct {
	Service
}

func NewUserService() *UserService {
	instance := new(UserService)
	return instance
}

func (service *UserService) CreateUser(user *model.User) (interface{}, error) {
	service.DataBase().Insert(user)
	return nil, nil
}
//func (service *UserService) GetUser() (interface{}, error) {
//	user := new(model.User)
//
//}
