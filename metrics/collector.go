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
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"tcloud_exporter/utils"
	"time"
)

type Tcloud_db interface {
	GetCode() string
	Rangeinstance(config *Config)
	AddInstance(request  *monitor.GetMonitorDataRequest, config *Config)
}

type MetricChannel struct {
	Apinamespace string
	MetricType   string
	InstanceName  Tcloud_db
	Config *Config
}

type Config struct{
	Mysql []Mysql_instance
	Redis []Redis_instance
	Mongodb []Mongodb_instance
	Kafka []Kafka_instance
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
				data := utils.GetMysqlMetrics(dataconfig)
				for _, mysqlmetric := range data {
					code := tclouddb.GetCode()
					metric_chan <- MetricChannel{Apinamespace: code, MetricType: mysqlmetric, InstanceName: tclouddb, Config: config}

				}
			case "mongodb":
				tclouddb = new(Mongodb)
				data := utils.GetMongoMetrics(dataconfig)
				for _, mongometric := range data {
					code := tclouddb.GetCode()
					metric_chan <- MetricChannel{Apinamespace: code, MetricType: mongometric, InstanceName: tclouddb, Config: config}
				}
			case "redis":
				tclouddb = new(Redis)
				data := utils.GetRedisMetrics(dataconfig)
				for _, mongometric := range data {
					code := tclouddb.GetCode()
					metric_chan <- MetricChannel{Apinamespace: code, MetricType: mongometric, InstanceName: tclouddb, Config: config}
				}
			}
		}
		time.Sleep(time.Second * 60)

	}
}

// 调度器，用于控制腾讯云接口访问频率
func Dispatch(id, key string, metric_chan chan MetricChannel, MetricCollector *MetricObj) {

	// 获取client
	client := GetClient(id, key)
	for {
		//value_temp := <- metric_chan
		for i := 0; i <= 10; i++ {
			fmt.Println("执行指标采集", i)
			GetMetrics(client, MetricCollector, <-metric_chan)
		}
		time.Sleep(time.Second * 1)
	}
}

func NamespaceToNameMap(namespace string) string {
	var name string
	switch namespace {
	case "QCE/CMONGO":
		name = "mongodb"
	case "QCE/CDB":
		name = "mysql"
	case "QCE/REDIS":
		name = "redis"
	default:
		name = "unknown"
	}
	return name

}

func NameToNamespaceMap(name string) string {
	var namespace string
	switch name {
	case "mysql":
		namespace = "QCE/CDB"
	case "mongodb":
		namespace = "QCE/CMONGO"
	case "redis":
		namespace = "QCE/REDIS"
	default:
		namespace = "unknown"
	}
	return namespace
}
