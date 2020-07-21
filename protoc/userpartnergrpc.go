package protoc

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql" //
	"xorm.io/xorm"
)

type UserPartnerSerice struct {
	db *xorm.Engine
}
//implemnet interface

func (c *UserPartnerSerice) connect() *UserPartnerSerice {
	engine, err := xorm.NewEngine("mysql", "root:123@tcp(0.0.0.0:7070)/exercise5")
	if err != nil {
		panic(err)
	}
	return &UserPartnerSerice{db: engine}

}

func (c *UserPartnerSerice) GetUserPartnerList(ctx context.Context, in *UserPartnerListRequest) (*UserPartnerListReponse, error) {
	up := UserPartnerSerice{}
	conn := up.connect()
	var listUser []*UserPartner

	err := conn.db.Find(&listUser)
	if err != nil {
		return nil, err
	}
	return &UserPartnerListReponse{ListUser: listUser}, nil
	//return data, nil
}

func (c *UserPartnerSerice) GetUserPartnerById(ctx context.Context, in *UserPartnerByIdRequest) (*UserPartnerByIdReponse, error) {
	up := UserPartnerSerice{}
	conn := up.connect()

	condition := &UserPartner{
		Id: in.Id,
	}
	ishas, err := conn.db.Get(condition)
	if err != nil {
		return nil, err
	}
	if !ishas {
		return nil, errors.New("Get by id failed")
	}
	return &UserPartnerByIdReponse{UsrPartner: condition}, nil
}

func (c *UserPartnerSerice) DeleteUserPartnerById(ctx context.Context, in *DeleteUserPartneRequest) (*DeleteUserPartnerReponse, error) {
	up := UserPartnerSerice{}
	conn := up.connect()
	rows, err := conn.db.Where("id=?", in.Id).Delete(&UserPartner{})
	if rows == 0 {
		return nil, errors.New("Delete Failed")
	}
	if err != nil {
		return nil, err
	}
	return &DeleteUserPartnerReponse{Status: "Success"}, nil

}

func (c *UserPartnerSerice) CreateUserPartner(ctx context.Context, in *CreateUserPartnerRequest) (*CreateUserPartnerReponse, error) {
	up := UserPartnerSerice{}
	conn := up.connect()
	rows, err := conn.db.Insert(in.UsrPartner)
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, errors.New("Insert failed")
	}
	return &CreateUserPartnerReponse{Status: "Success"}, nil

}
