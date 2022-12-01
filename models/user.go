package models

type User struct {
	// id
	Id int `json:"id" swaggerignore:"true"`
	// 名称
	Name string `json:"name"`
	// 密码
	Password string `json:"pswd" gorm:"column:pswd"`
	// 线上状态消息
	StateMsg string `json:"stateMsg" gorm:"column:stateMsg"`
	// 头像地址
	Portrait string `json:"portrait"`
	// 消息id
	Mid string `json:"mid"`
}

// 给操作的表 取别名, 新版本必须这么做
func (u *User) TableName() string {
	return "user"
}
