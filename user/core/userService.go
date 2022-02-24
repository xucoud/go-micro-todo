package core

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/errors"
	"go_code/go-micro-todo/user/model"
	"go_code/go-micro-todo/user/services"
)

func BuildUser(user model.User) *services.UserModel {
	return &services.UserModel{
		ID: uint32(user.ID),
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
		UpdateAt: user.UpdatedAt.Unix(),
	}
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, rep *services.UserResponse) error {
	var user model.User
	rep.Code = 200
	if err := model.Db.Table("users").Where("user_name = ?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			rep.Code = 400
			return nil
		}
		rep.Code = 500
		return nil
	}
	if user.CheckPassword(req.Password) == false {
		rep.Code = 400
		return nil
	}
	rep.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, rep *services.UserResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New( "", "两次密码不一致！", 400)
		return err
	}
	count := 0
	if err := model.Db.Table("users").Where("user_name = ?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("", "该用户已经存在！", 400)
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.Db.Create(&user).Error; err != nil {
		return err
	}
	rep.UserDetail = BuildUser(user)
	return nil
}
