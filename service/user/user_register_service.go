package user

import (
	"context"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
)

type UserRegisterService struct {
	NickName string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName string `form:"username" json:"username" binding:"required,min=3,max=80"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=40"`
	Code     string `form:"code" json:"code" binding:"required,min=6,max=6"`
}

type registerResponse struct {
	Token string `json:"token"`
	serializer.User
}

// vaild check if regiser info correct
func (service *UserRegisterService) vaild() *serializer.Response {
	if service.Code != cache.RedisClient.Get(context.Background(), cache.EmailCodeKey(service.UserName)).Val() {
		return &serializer.Response{
			Code: serializer.CodeParamsError,
			Msg:  "Code can't matched",
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

	// signed jwt token
	token, err := utils.GenToken("miaoqiu", 24, &user)
	if err != nil {
		return serializer.Err(serializer.CodeError, "token generate error", err)
	}

	return serializer.Success(registerResponse{
		Token: token,
		User:  serializer.BuildUser(user),
	})
}
