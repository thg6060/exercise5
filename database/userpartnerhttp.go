package database

import (
	"errors"

	"github.com/thg6060/exercise5/protoc"
)

type UserPartApi struct {
}

func (up *UserPartApi) GetUserByCdt(condition *protoc.UserPartnerByIdRequest) ([]*protoc.UserPartner, error) {
	var listUser []*protoc.UserPartner
	temp := protoc.UserPartnerByIdRequest{
		UserId: condition.UserId,
		Phone:  condition.Phone,
	}
	err := db.Limit(int(condition.GetLimit()), 0).Find(&listUser, temp)
	if err != nil {
		return nil, err
	}
	return listUser, nil
}
func (up *UserPartApi) GetUserList() ([]*protoc.UserPartner, error) {
	var listUser []*protoc.UserPartner

	err := db.Find(&listUser)
	if err != nil {
		return nil, err
	}
	return listUser, nil
}

func (up *UserPartApi) DeleteUser(id string) error {

	rows, err := db.Where("id=?",id).Delete(&protoc.UserPartner{})
	if rows == 0 {
		return errors.New("Delete Failed")
	}
	if err != nil {
		return err
	}
	return nil
}

func (up *UserPartApi) GetUserById(id string) (*protoc.UserPartner, error) {
	//var User *protoc.UserPartner
	condition := &protoc.UserPartner{
		Id: id,	
	}
	ishas,err := db.Get(condition)
	if err != nil {
		return nil, err
	}
	if !ishas{
		return nil,errors.New("Get by id failed")
	}
	return condition, nil
}
func (up *UserPartApi) CreateUser(usr *protoc.UserPartner) error {
	//var User *protoc.UserPartner
	rows,err:=db.Insert(usr)
	if err!=nil{
		return err
	}
	if rows==0{
		return errors.New("Insert failed")
	}
	return nil
}


