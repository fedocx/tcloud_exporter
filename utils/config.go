package utils

import "github.com/spf13/viper"

func GetAuthInfo(resourceconfig *viper.Viper) (string, string) {
	TENCENTCLOUD_SECRET_ID := resourceconfig.GetString("tencentcloud_secret_id")
	TENCENTCLOUD_SECRET_KEY := resourceconfig.GetString("tencentcloud_secret_key")
	return TENCENTCLOUD_SECRET_ID, TENCENTCLOUD_SECRET_KEY
}

func GetMysqlInstance(resourceconfig *viper.Viper) map[string][]string {
	//mysql := resourceconfig.GetStringSlice("mysql")
	mysql := resourceconfig.GetStringMapStringSlice("mysql")
	return mysql
}

func GetMongoInstance(resourceconfig *viper.Viper) map[string][]string {
	mongo := resourceconfig.GetStringMapStringSlice("mongodb")
	return mongo
}

func GetRedisInstance(resourceconfig *viper.Viper) map[string]string {
	redis := resourceconfig.GetStringMapString("redis")
	return redis
}

func GetMysqlMetrics(dataconfig *viper.Viper) []string {
	mysql := dataconfig.GetStringSlice("mysql")
	return mysql
}

func GetMongoMetrics(dataconfig *viper.Viper) []string {
	mongo := dataconfig.GetStringSlice("mongodb")
	return mongo
}

func GetMongoClusterMetrics(dataconfig *viper.Viper) []string {
	mongo_cluster := dataconfig.GetStringSlice("mongodb_cluster")
	return mongo_cluster
}

func GetMongoReplicationMetrics(dataconfig *viper.Viper) []string {
	mongo_cluster := dataconfig.GetStringSlice("mongodb_replication")
	return mongo_cluster
}
func GetRedisMetrics(dataconfig *viper.Viper) []string {
	redis := dataconfig.GetStringSlice("redis")
	return redis
}

func GetKfakaTopicMetrics(dataconfig *viper.Viper) []string {
	redis := dataconfig.GetStringSlice("kafka_topic")
	return redis
}
func GetKfakaPartitionMetrics(dataconfig *viper.Viper) []string {
	redis := dataconfig.GetStringSlice("kafka_partition")
	return redis
}
