// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package metrics

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Tcloud_db interface {
	GetCode() string
	GetMetrics(dataconfig *viper.Viper) []string
}

type MetricChannel struct {
	Apinamespace string
	MetricType   string
	InstanceName  Tcloud_db
	Config []map[string]string
	Type string
}

type Config struct{
	Mysql []map[string]string `mapstructure:"mysql"`
	Redis []map[string]string `mapstructure:"redis"`
	Mongodb []map[string]string `mapstructure:"mongodb"`
	KafkaTopic []map[string]string `mapstructure:"kafka_topic"`
	KafkaPartition []map[string]string `mapstructure:"kafka_partition"`

}

// 根据当前配置信息，获取配置里面的数据库项，并根据数据库项获取响应的数据库指标,通过goroutin方式执行
func GetResourceList(resourceconfig *viper.Viper, dataconfig *viper.Viper, metric_chan chan MetricChannel) {

	// 获取配置文件采集项
	objects := dataconfig.AllKeys()
	config  := new(Config)
	err := resourceconfig.Unmarshal(&config)
	if err !=nil{
		panic(err)
	}
	for {
		for _, val := range objects {
			var tclouddb Tcloud_db
			switch val {
			case "mysql":
				tclouddb = new(Mysql)
				//instancelist := utils.GetMysqlInstance(resourceconfig)
				data := tclouddb.GetMetrics(dataconfig)
				for _, mysqlmetric := range data {
					code := tclouddb.GetCode()
					mysql_config := config.Mysql
					metric_chan <- MetricChannel{Apinamespace: code, MetricType: mysqlmetric, InstanceName: tclouddb, Config: mysql_config,Type: val}

				}
			case "mongodb":
				tclouddb = new(Mongodb)
				data := tclouddb.GetMetrics(dataconfig)
				for _, mongometric := range data {
					code := tclouddb.GetCode()
					mongodb_config := config.Mongodb
					metric_chan <- MetricChannel{Apinamespace: code, MetricType: mongometric, InstanceName: tclouddb, Config: mongodb_config,Type: val}
				}
			case "redis":
				tclouddb = new(Redis)
				data := tclouddb.GetMetrics(dataconfig)
				for _, mongometric := range data {
					code := tclouddb.GetCode()
					redis_config := config.Redis
					metric_chan <- MetricChannel{Apinamespace: code, MetricType: mongometric, InstanceName: tclouddb, Config: redis_config,Type: val}
				}
			case "kafka_topic":
				tclouddb = new(Kafka_topic)
				data := tclouddb.GetMetrics(dataconfig)
				for _, mongometric := range data {
					code := tclouddb.GetCode()
					kafka_config := config.KafkaTopic
					//metric_chan <- MetricChannel{Apinamespace: code, MetricType: mongometric, InstanceName: tclouddb, Config: kafka_config}
					metric_chan <- MetricChannel{Apinamespace: code, MetricType: mongometric, InstanceName: tclouddb, Config: kafka_config,Type: val}
				}
			case "kafka_partition":
				tclouddb = new(Kafka_partition)
				data := tclouddb.GetMetrics(dataconfig)
				for _, mongometric := range data {
					code := tclouddb.GetCode()
					kafka_config := config.KafkaPartition
					metric_chan <- MetricChannel{Apinamespace: code, MetricType: mongometric, InstanceName: tclouddb, Config: kafka_config,Type: val}
				}
			}
		}
		time.Sleep(time.Second * 60)

	}
}

// 调度器，用于控制腾讯云接口访问频率
func Dispatch(id, key string, metric_chan chan MetricChannel, MetricCollector *MetricObj) {
	lock := make(chan int,10)
	// init lock
	for i:=1; i< 11; i ++ {
		lock <- i
	}

	// 获取client
	client := GetClient(id, key)
	for {
		//value_temp := <- metric_chan
		for i := 0; i <= 10; i++ {
			fmt.Println("执行指标采集", i)
			go GetMetrics(client, MetricCollector, <-metric_chan,lock)
		}
		time.Sleep(time.Second * 2)
	}
}

//func NamespaceToNameMap(namespace string) string {
//	var name string
//	switch namespace {
//	case "QCE/CMONGO":
//		name = "mongodb"
//	case "QCE/CDB":
//		name = "mysql"
//	case "QCE/REDIS":
//		name = "redis"
//	default:
//		name = "unknown"
//	}
//	return name
//
//}
//
//func NameToNamespaceMap(name string) string {
//	var namespace string
//	switch name {
//	case "mysql":
//		namespace = "QCE/CDB"
//	case "mongodb":
//		namespace = "QCE/CMONGO"
//	case "redis":
//		namespace = "QCE/REDIS"
//	default:
//		namespace = "unknown"
//	}
//	return namespace
//}
