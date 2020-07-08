package utils

import "github.com/spf13/viper"

func GetAuthInfo()(string,string){
	TENCENTCLOUD_SECRET_ID := viper.GetString("tencentcloud_secret_id")
	TENCENTCLOUD_SECRET_KEY := viper.GetString("tencentcloud_secret_key")
	return TENCENTCLOUD_SECRET_ID,TENCENTCLOUD_SECRET_KEY
}

func GetMysqlInstance()[]string{
	mysql := viper.GetStringSlice("mysql")
	return  mysql
}