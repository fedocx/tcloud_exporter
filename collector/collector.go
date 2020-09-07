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
package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"log"
	"tcloud_exporter/metrics"
	"tcloud_exporter/utils"
	"time"
)

type Metric_Gauge struct{
	Gauge map[string]*prometheus.GaugeVec
}
type Object_Gauge struct{
	Object map[string][]*Metric_Gauge
}

// 消费指标
func MetricsConsumer(metrics *metrics.MetricObj, dataconfig *viper.Viper,flush_metrics chan bool){
	gauge_maps := make(map[string][]*Metric_Gauge)
	object_gauge := new(Object_Gauge)
	object_gauge.Object = gauge_maps
	MysqlRegister(metrics,dataconfig,object_gauge)
	MongodbRegister(metrics,dataconfig,object_gauge)
	RedisRegister(metrics,dataconfig,object_gauge)
	KafkaPartitionRegister(metrics,dataconfig,object_gauge)
	KafkaTopicRegister(metrics,dataconfig,object_gauge)

	if <-flush_metrics{
		ReadMetrics(metrics,dataconfig,object_gauge)

	}

}

// 定义监控指标，对于mysql采集哪些指标，并对这些指标进行注册。
func Register(namespace, subsystem, name, help string)(*prometheus.GaugeVec){
	a := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name: name,
		Help: help,
	}, []string{"instance"})
	prometheus.MustRegister(a)
	return a
}

// 获取mysql的指标信息
func MysqlRegister(metriccollector *metrics.MetricObj,dataconfig *viper.Viper,object_gauge *Object_Gauge){
	gauge_map := make(map[string]*prometheus.GaugeVec)
	gauge_maps := new(Metric_Gauge)
	gauge_maps.Gauge = gauge_map
	mysql_metrics := utils.GetMysqlMetrics(dataconfig)
	for _,val := range mysql_metrics{
		gauge_maps.Gauge[val] = Register("tcloud","database_mysql",val,"Number of blob storage operations wa=itingto be processed.")
	}
	object_gauge.Object["mysql"] = append(object_gauge.Object["mysql"],gauge_maps)
}

func MongodbRegister(metriccollector *metrics.MetricObj,dataconfig *viper.Viper,object_gauge *Object_Gauge){
	gauge_map := make(map[string]*prometheus.GaugeVec)
	gauge_maps := new(Metric_Gauge)
	gauge_maps.Gauge = gauge_map

	mongodb_metrics := utils.GetMongoMetrics(dataconfig)
	for _,val := range mongodb_metrics{
		gauge_maps.Gauge[val] = Register("tcloud","database_mongodb",val,"Number of blob storage operations wa=itingto be processed.")

	}
	object_gauge.Object["mongodb"] = append(object_gauge.Object["mongodb"],gauge_maps)
}

func RedisRegister(metriccollector *metrics.MetricObj,dataconfig *viper.Viper,object_gauge *Object_Gauge){
	gauge_map := make(map[string]*prometheus.GaugeVec)
	gauge_maps := new(Metric_Gauge)
	gauge_maps.Gauge = gauge_map

	mongodb_metrics := utils.GetRedisMetrics(dataconfig)
	for _,val := range mongodb_metrics{
		gauge_maps.Gauge[val] = Register("tcloud","database_redis",val,"Number of blob storage operations wa=itingto be processed.")
		log.Print("redis注册成功",val)

	}
	object_gauge.Object["redis"] = append(object_gauge.Object["redis"],gauge_maps)
}

func KafkaTopicRegister(metriccollector *metrics.MetricObj,dataconfig *viper.Viper,object_gauge *Object_Gauge){
	gauge_map := make(map[string]*prometheus.GaugeVec)
	gauge_maps := new(Metric_Gauge)
	gauge_maps.Gauge = gauge_map

	mongodb_metrics := utils.GetKfakaTopicMetrics(dataconfig)
	for _,val := range mongodb_metrics{
		gauge_maps.Gauge[val] = Register("tcloud","database_kafka_topic",val,"Number of blob storage operations wa=itingto be processed.")
		log.Print("kafka topic注册成功",val)

	}
	object_gauge.Object["kafka_topic"] = append(object_gauge.Object["kafka_topic"],gauge_maps)
}

func KafkaPartitionRegister(metriccollector *metrics.MetricObj,dataconfig *viper.Viper,object_gauge *Object_Gauge){
	gauge_map := make(map[string]*prometheus.GaugeVec)
	gauge_maps := new(Metric_Gauge)
	gauge_maps.Gauge = gauge_map

	mongodb_metrics := utils.GetKfakaPartitionMetrics(dataconfig)
	for _,val := range mongodb_metrics{
		gauge_maps.Gauge[val] = Register("tcloud","database_kafka_partition",val,"Number of blob storage operations wa=itingto be processed.")
		log.Print("kafka partition注册成功",val)

	}
	object_gauge.Object["kafka_partition"] = append(object_gauge.Object["kafka_partition"],gauge_maps)
}
func ReadMetrics(metriccollector *metrics.MetricObj,dataconfig *viper.Viper,object_gauge *Object_Gauge){

	// 休息10s钟，等接口从腾讯云获取完数据,存放到metriccollector之后再开始
	time.Sleep(time.Second * 10)
	// 获取指标
	//namespace := metrics.GetMysqlCode()
	//name := metrics.NamespaceToNameMap(namespace)
	// key mysql
	for {
		for key,val := range object_gauge.Object{
			log.Print("再次更新数据",key)
			for _,vec := range val{
				// metric_name : netin   metric_value 123
				for metric_name,metric_value := range vec.Gauge{
					GetGuage(metric_name,metric_value,metriccollector,key)

				}
			}
		}
		time.Sleep(time.Second * 60)

	}

}