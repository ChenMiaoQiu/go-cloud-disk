package serializer

import "github.com/ChenMiaoQiu/go-cloud-disk/model"

// User serializer
type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser return a user serializer
func BuildUser(user model.User) User {
	return User{
		ID:       user.Uuid,
		UserName: user.UserName,
		NickName: user.NickName,
		Status:   user.Status,
		Avatar:   user.Avatar,
	}
}

// BuildUsers return user serializers
func BuildUsers(users []model.User) (usersSerializer []User) {
	for _, user := range users {
		usersSerializer = append(usersSerializer, BuildUser(user))
	}
	return
}
