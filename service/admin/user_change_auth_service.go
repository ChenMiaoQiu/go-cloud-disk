package admin

import (
	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/ChenMiaoQiu/go-cloud-disk/serializer"
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
		return serializer.DBErr("get user info err when change user auth", err)
	}

	if user.Uuid == "" {
		return serializer.ParamsErr("can't get user when change user auth", nil)
	}

	// check if user is an admin
	if userStatus != model.StatusAdmin && userStatus != model.StatusSuperAdmin {
		return serializer.NotAuthErr("common user can't change auth")
	}

	// normal admin can't change admin auth
	if userStatus == model.StatusAdmin {
		if user.Status == model.StatusAdmin || user.Status == model.StatusSuperAdmin {
			return serializer.NotAuthErr("admin not auth to change admin auth")
		}

		if service.NewStatus == model.StatusAdmin || service.NewStatus == model.StatusSuperAdmin {
			return serializer.NotAuthErr("admin not auth to change admin auth")
		}
	}

	// save user auth
	user.Status = service.NewStatus
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.DBErr("save user info err when change user auth", err)
	}
	return serializer.Success(serializer.BuildUser(user))
}
