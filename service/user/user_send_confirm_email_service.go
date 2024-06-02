package user

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
)

type UserSendConfirmEmailService struct {
	UserEmail string `json:"email" form:"email" binding:"required"`
}

func getConfirmCode() string {
	var confirmCode int
	for i := 0; i < 6; i++ {
		confirmCode = confirmCode*10 + (rand.Intn(9) + 1)
	}
	confirmCodeStr := strconv.Itoa(confirmCode)
	return confirmCodeStr
}

func (service *UserSendConfirmEmailService) SendConfirmEmail() serializer.Response {
	// check email format
	if !utils.VerifyEmailFormat(service.UserEmail) {
		return serializer.ParamsErr("email format err when send confirm email", nil)
	}
	// check user request email times in recent
	if cache.RedisClient.Get(context.Background(), cache.RecentSendUserKey(service.UserEmail)).Val() != "" {
		return serializer.ParamsErr("multi request send email", nil)
	}

	// check if email has register
	var emailNum int64
	if err := model.DB.Model(&model.User{}).Where("user_name = ?", service.UserEmail).Count(&emailNum).Error; err != nil {
		return serializer.DBErr("get user err", err)
	}
	if emailNum > 0 {
		return serializer.ParamsErr("has register", nil)
	}

	code := getConfirmCode()
	cache.RedisClient.Set(context.Background(), cache.EmailCodeKey(service.UserEmail), code, time.Minute*30)

	if err := utils.SendConfirmMessage(service.UserEmail, code); err != nil {
		return serializer.DBErr("send email err", err)
	}
	// limit 1 email max request 1 confirm email in 3 minute
	cache.RedisClient.Set(context.Background(), cache.RecentSendUserKey(service.UserEmail), code, time.Minute*3)

	return serializer.Success(nil)
}
