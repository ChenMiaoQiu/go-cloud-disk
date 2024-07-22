package user

import (
	"context"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/rabbitMQ"
	"github.com/ChenMiaoQiu/go-cloud-disk/rabbitMQ/task"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils"
	"github.com/ChenMiaoQiu/go-cloud-disk/utils/logger"
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
		return serializer.ParamsErr("NotEmail", nil)
	}
	// check user request email times in recent
	if cache.RedisClient.Get(context.Background(), cache.RecentSendUserKey(service.UserEmail)).Val() != "" {
		return serializer.ParamsErr("HasSendCode", nil)
	}

	// check if email has register
	var emailNum int64
	if err := model.DB.Model(&model.User{}).Where("user_name = ?", service.UserEmail).Count(&emailNum).Error; err != nil {
		logger.Log().Error("[UserSendConfirmEmailService.SendConfirmEmail] Fail to find user: ", err)
		return serializer.DBErr("", err)
	}
	if emailNum > 0 {
		return serializer.ParamsErr("HasRegister", nil)
	}

	code := getConfirmCode()
	cache.RedisClient.Set(context.Background(), cache.EmailCodeKey(service.UserEmail), code, time.Minute*30)

	if err := service.sendConfirmEmailToMQ(service.UserEmail, code); err != nil {
		return serializer.InternalErr("", err)
	}
	// limit 1 email max request 1 confirm email in 3 minute
	cache.RedisClient.Set(context.Background(), cache.RecentSendUserKey(service.UserEmail), code, time.Minute*3)

	return serializer.Success(nil)
}

func (service *UserSendConfirmEmailService) sendConfirmEmailToMQ(targetEmail string, code string) error {
	// limit 1 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	sendConfirmEmailReq := task.SendConfirmEmailRequest{
		Email: targetEmail,
		Code:  code,
	}

	body, err := json.Marshal(sendConfirmEmailReq)
	if err != nil {
		logger.Log().Error("[UserSendConfirmEmailService.SendConfirmEmailToMQ] Fail to marshal request: ", err)
		return err
	}
	err = rabbitMQ.SendMessageToMQ(ctx, rabbitMQ.RabbitMqSendEmailQueue, body)
	if err != nil {
		return err
	}

	return nil
}
