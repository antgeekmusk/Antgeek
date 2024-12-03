package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
func InitConfig()  {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("====== config app init ======")
	fmt.Println("config.mysql : ",viper.Get("mysql"))
}

func InitMysql()  {
	var err error
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",viper.GetString("mysql.username"),viper.GetString("mysql.password"),viper.GetString("mysql.ip"),viper.GetString("mysql.port"),viper.GetString("mysql.db"),viper.GetString("mysql.conn_params"))
	Db, err = gorm.Open(mysql.Open(connStr),&gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("====== InitMysql init ======")
}