package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name string
	Password string
	Phone string
	Email string
	Identity string
	ClientIp string
	ClientPort string
	LoginTime time.Time
	HeartbeatTime time.Time
	LogoutTime time.Time `gorm:"column:logout_time"`
	IsLogout bool
	DeviceInfo string
}

func (table UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic,10)
	utils.Db.Find(&data)
	return data
}