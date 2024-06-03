package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
	loglog "github.com/ChenMiaoQiu/go-cloud-disk/utils/log"
)

type UserChangeAuthService struct {
	UserId    string `json:"userid" form:"userid" required:"binding"`
	NewStatus string `json:"status" form:"status" required:"binding"`
}

// UserChangeAuth change user auth need input user status that use this func
func (service *UserChangeAuthService) UserChangeAuth(userStatus string) serializer.Response {
	// get user info from database
	var user model.User
	if err := model.DB.Where("uuid = ?", service.UserId).Find(&user).Error; err != nil {
		loglog.Log().Error("[UserChangeAuthService.UserChangeAuth] Fail to find user info: ", err)
		return serializer.DBErr("", err)
	}

	if user.Uuid == "" {
		return serializer.ParamsErr("", nil)
	}

	// check if user is an admin
	if userStatus != model.StatusAdmin && userStatus != model.StatusSuperAdmin {
		return serializer.NotAuthErr("")
	}

	// normal admin can't change admin auth
	if userStatus == model.StatusAdmin {
		if user.Status == model.StatusAdmin || user.Status == model.StatusSuperAdmin {
			return serializer.NotAuthErr("")
		}

		if service.NewStatus == model.StatusAdmin || service.NewStatus == model.StatusSuperAdmin {
			return serializer.NotAuthErr("")
		}
	}

	// save user auth
	user.Status = service.NewStatus
	if err := model.DB.Save(&user).Error; err != nil {
		loglog.Log().Error("[UserChangeAuthService.UserChangeAuth] Fail to save user info: ", err)
		return serializer.DBErr("", err)
	}
	return serializer.Success(serializer.BuildUser(user))
}
