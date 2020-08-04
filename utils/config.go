package utils

import "github.com/spf13/viper"

func GetAuthInfo(resourceconfig *viper.Viper)(string,string){
	TENCENTCLOUD_SECRET_ID := resourceconfig.GetString("tencentcloud_secret_id")
	TENCENTCLOUD_SECRET_KEY := resourceconfig.GetString("tencentcloud_secret_key")
	return TENCENTCLOUD_SECRET_ID,TENCENTCLOUD_SECRET_KEY
}

func GetMysqlInstance(resourceconfig *viper.Viper)[]string{
	mysql := resourceconfig.GetStringSlice("mysql")
	return  mysql
}

func GetMongoInstance(resourceconfig *viper.Viper)[]string{
	mongo := resourceconfig.GetStringSlice("mongodb")
	return  mongo
}
func GetMysqlMetrics(dataconfig *viper.Viper)[]string{
	mysql := dataconfig.GetStringSlice("mysql")
	return mysql
}

func GetMongoMetrics(dataconfig *viper.Viper)[]string{
	mongo := dataconfig.GetStringSlice("mongodb")
	return mongo
}
