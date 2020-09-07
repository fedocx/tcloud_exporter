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
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"tcloud_exporter/collector"
	"tcloud_exporter/metrics"
	"tcloud_exporter/utils"
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
	var flush_metrics = make(chan bool)


	go metrics.GetResourceList(resourceconfig,dataconfig, metric_channel,flush_metrics)
	// 消费指标，从腾讯云获取监控数据，并将指标存放到MetricCollector列表中
	go metrics.Dispatch(TENCENTCLOUD_SECRET_ID, TENCENTCLOUD_SECRET_KEY,  metric_channel, MetricCollector)

	go collector.MetricsConsumer(MetricCollector,dataconfig,flush_metrics)


	flag.Parse()
	http.Handle("/metrics",promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr,nil))
	//time.Sleep(time.Minute * 2)
	//resourceconfig,_ := InitConfig()
	//fmt.Println(utils.GetMysqlInstance(resourceconfig))
	//fmt.Println(resourceconfig.Get("mysql"))

}
