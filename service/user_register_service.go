package service

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
)

type UserRegisterService struct {
	NickName        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// vaild check if regiser info correct
func (service *UserRegisterService) vaild() *serializer.Response {
	// check password
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: serializer.CodeParamsError,
			Msg:  "Entered passwords differ",
		}
	}

	// check nickname
	count := int64(0)
	model.DB.Model(&model.User{}).Where("nick_name = ?", service.NickName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: serializer.CodeParamsError,
			Msg:  "Nickname occupied",
		}
	}

	// check username
	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "Username occupied",
		}
	}

	return nil
}

// Register check if register info correct. if it correct,
// register the user to database. Otherwise, return a error message
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model.StatusActiveUser,
	}

	// check user vaild
	if err := service.vaild(); err != nil {
		return *err
	}

	// encryption password
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(serializer.CodeError, "encrypation password err", err)
	}

	// create user
	if err := user.CreateUser(); err != nil {
		return serializer.ParamsErr("create user error", err)
	}

	return serializer.Success(serializer.BuildUser(user))
}
