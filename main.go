package main

import (
	"fmt"
	"im-golang/dao"
	"im-golang/server"
	"os"

	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := dao.InitDB()
	defer db.Close()

	r := server.NewRouter()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("配置读取错误")
	}
}
