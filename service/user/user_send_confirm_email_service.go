package user

import (
	"context"
	"math/rand"
	"net/mail"
	"strconv"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
)

type UserSendConfirmEmailService struct {
	UserEmail string `json:"email" form:"email" binding:"required"`
}

func getConfirmCode() string {
	var confirmCode int
	for i := 0; i < 6; i++ {
		confirmCode = confirmCode*10 + rand.Intn(10)
	}
	confirmCodeStr := strconv.Itoa(confirmCode)
	return confirmCodeStr
}

func (service *UserSendConfirmEmailService) SendConfirmEmail() serializer.Response {
	// check email format
	userEmail, err := mail.ParseAddress(service.UserEmail)
	if err != nil {
		return serializer.ParamsErr("email err when send confirm email", err)
	}
	// check user request email times in recent
	if cache.RedisClient.Get(context.Background(), cache.RecentSendUserKey(userEmail.Address)).Val() != "" {
		return serializer.ParamsErr("multi request send email", nil)
	}

	code := getConfirmCode()
	cache.RedisClient.Set(context.Background(), cache.EmailCodeKey(userEmail.Address), code, time.Minute*30)

	if err := utils.SendConfirmMessage(service.UserEmail, code); err != nil {
		return serializer.DBErr("send email err", err)
	}
	// limit 1 email max request 1 confirm email in 3 minute
	cache.RedisClient.Set(context.Background(), cache.RecentSendUserKey(userEmail.Address), code, time.Minute*3)

	return serializer.Success(nil)
}
