package servicepackage

import "qcengine/src/model"

type UserService struct {
	Service
}

func NewUserService() *UserService {
	instance := new(UserService)
	return instance
}

func (service *UserService) CreateUser(user *model.User) (interface{}, error) {
	res, err := service.DataBase().Insert(user)
	return res, err
}
func (service *UserService) DeleteUser(user *model.User) error {
	err := service.DataBase().Delete(user)
	return err
}
func (service *Service) UpdateUser(user *model.User) error {
	err := service.DataBase().Update(user)
	return err
}
func (service *Service) FindUser(user *model.User) (interface{}, error) {
	res, err := service.DataBase().Find(user)
	return res, err
}
//func (service *UserService) GetUser() (interface{}, error) {
//	user := new(model.User)
//
//}
