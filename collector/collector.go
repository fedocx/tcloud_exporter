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
	"tcloud_exporter/metrics"
	"time"
)

// 消费指标
func MetricsConsumer(metrics *metrics.MetricObj){
	go MysqlRegister(metrics)
	go MongodbRegister(metrics)
}

// 定义监控指标，对于mysql采集哪些指标，并对这些指标进行注册。
func Register(namespace, subsystem, name, help string)(*prometheus.GaugeVec){
	a := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name: name,
		Help: help,
	}, []string{"instance"})
	return a
}

// 获取mysql的指标信息
func MysqlRegister(metriccollector *metrics.MetricObj){
	namespace := metrics.GetMysqlCode()
	name := metrics.NamespaceToNameMap(namespace)
	metrics.GetMysqlCode()
	mysqlDiskUse := Register("tcloud","database_mysql","disk","Number of blob storage operations wa=itingto be processed.")
	mysqlCpuUseRate := Register("tcloud","database_mysql","cpu","Number of blob storage operations wa=itingto be processed.")
	mysqlMemUseRate := Register("tcloud","database_mysql","mem","Number of blob storage operations wa=itingto be processed.")
	mysqlNetOut := Register("tcloud","database_mysql","net_in","Number of blob storage operations wa=itingto be processed.")
	mysqlNetIn := Register("tcloud","database_mysql","net_out","Number of blob storage operations wa=itingto be processed.")

	// 注册
	prometheus.MustRegister(mysqlDiskUse)
	prometheus.MustRegister(mysqlCpuUseRate)
	prometheus.MustRegister(mysqlMemUseRate)
	prometheus.MustRegister(mysqlNetOut)
	prometheus.MustRegister(mysqlNetIn)

	// 获取指标
	for {
		GetGuage("CPUUseRate",mysqlCpuUseRate,metriccollector,name)
		GetGuage("MemoryUseRate",mysqlMemUseRate,metriccollector,name)
		GetGuage("VolumeRate",mysqlDiskUse,metriccollector,name)
		GetGuage("BytesSent",mysqlNetOut,metriccollector,name)
		GetGuage("BytesReceived",mysqlNetIn,metriccollector,name)
		time.Sleep(time.Second * 60)

	}

}

func MongodbRegister(metriccollector *metrics.MetricObj){
	namespace := metrics.GetMongoCode()
	name := metrics.NamespaceToNameMap(namespace)
	metrics.GetMysqlCode()
	mongoNetin := Register("tcloud","database_mongo","netin","Number of blob storage operations wa=itingto be processed.")
	mongoNetout := Register("tcloud","database_mongo","netout","Number of blob storage operations wa=itingto be processed.")

	// 注册
	prometheus.MustRegister(mongoNetin)
	prometheus.MustRegister(mongoNetout)

	// 获取指标
	for {
		GetGuage("Netin", mongoNetin, metriccollector, name)
		GetGuage("Netout", mongoNetout, metriccollector, name)
		time.Sleep(time.Second * 60)
	}

}