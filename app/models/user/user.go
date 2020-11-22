package user

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordComfirm string `gorm:"-" valid:"password_comfirm"`
}

// ComparePassword 对比密码是否匹配
func (u User) ComparePassword(_password string) bool {
	// return password.CheckHash(_password, u.Password)
	return u.Password == _password
}

// Link 方法用来生成用户链接
func (u User) Link() string {
	return route.Name2URL("users.show", "id", u.GetStringID())
}
