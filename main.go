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
package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"tcloud_exporter/collector"
	"tcloud_exporter/metrics"
	"tcloud_exporter/utils"
	"time"
)

var TENCENTCLOUD_SECRET_ID, TENCENTCLOUD_SECRET_KEY string

func InitConfig() (resourceconfig, dataconfig *viper.Viper) {
	workDir, _ := os.Getwd()

	resourceconfig = viper.New()
	resourceconfig.SetConfigName("tencent")
	resourceconfig.SetConfigType("yml")
	resourceconfig.AddConfigPath(workDir + "/config")
	err := resourceconfig.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dataconfig = viper.New()
	dataconfig.SetConfigName("metrics")
	dataconfig.SetConfigType("yml")
	dataconfig.AddConfigPath(workDir + "/config")
	err = dataconfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return resourceconfig, dataconfig
}

var addr = flag.String("listen-addr", ":8081", "the port to listen on for HTTP requests")


func main() {

	resourceconfig, dataconfig := InitConfig()
	TENCENTCLOUD_SECRET_ID, TENCENTCLOUD_SECRET_KEY = utils.GetAuthInfo(resourceconfig)

	// 定义collector 并通过collector来采集数据
	MetricCollector := new(metrics.MetricObj)
	MetricCollector.Products = make(map[string][]*metrics.Product)

	// 生成channel，用于对指标进行消费
	var metric_channel = make(chan metrics.MetricChannel,10)


	go metrics.GetResourceList(resourceconfig,dataconfig, metric_channel)
	// 消费指标，从腾讯云获取监控数据，并将指标存放到MetricCollector列表中
	go metrics.Dispatch(TENCENTCLOUD_SECRET_ID, TENCENTCLOUD_SECRET_KEY,  metric_channel, MetricCollector)

	go collector.MetricsConsumer(MetricCollector)


	// 首先调用各个数据库的采集接口，获取到采集指标。
	// 其次在采集接口中通过注册的方式控制获取哪些指标，比如disk或者net的指标
	// 将获取到的指标放入到collector中，api接口通过gorouting方式执行，定时按照特定频率更新指标信息到collector中
	// collector消费端通过gorouting方式从collector中获取指标信息，完成对collecotor的消费工作。

	//将指标添加到指标库中,针对gauge类型的指标
	//go func(){
	//	for {
	//		// 通过调用腾讯云接口，按照指定频率获取监控指标，生产指标
	//		mysqlmetrics := metrics.GetMysqlMetrics(TENCENTCLOUD_SECRET_ID,TENCENTCLOUD_SECRET_KEY,mysqlmetrics)
	//
	//		// 通过获取的指标，存放在mysqlmetrics 变量中，并将指标传递给指标消费函数，该函数获取对应的值，在Prometheus中进行展示
	//		collector.MetricsConsumer(mysqlmetrics)
	//
	//		// 每隔1分钟获取一次指标
	//		time.Sleep(time.Second * 60)
	//
	//	}
	//}()
	//
	flag.Parse()
	http.Handle("/metrics",promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr,nil))
	time.Sleep(time.Minute * 2)

}
