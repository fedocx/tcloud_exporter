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
)

// 消费指标
func MetricsConsumer(metrics *metrics.MetricObj){
	MysqlRegister(metrics)
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
func MysqlRegister(metrics *metrics.MetricObj){
	mysqlDiskUse := Register("tcloud","database_mysql","disk","Number of blob storage operations wa=itingto be processed.")
	mysqlCpuUseRate := Register("tcloud","database_mysql","disk","Number of blob storage operations wa=itingto be processed.")
	mysqlMemUseRate := Register("tcloud","database_mysql","disk","Number of blob storage operations wa=itingto be processed.")
	mysqlNetOut := Register("tcloud","database_mysql","disk","Number of blob storage operations wa=itingto be processed.")
	mysqlNetIn := Register("tcloud","database_mysql","disk","Number of blob storage operations wa=itingto be processed.")

	GetGuage("CPUUseRate",mysqlCpuUseRate,metrics)
	GetGuage("MemoryUseRate",mysqlMemUseRate,metrics)
	GetGuage("VolumeRate",mysqlDiskUse,metrics)
	GetGuage("BytesSent",mysqlNetOut,metrics)
	GetGuage("BytesReceived",mysqlNetIn,metrics)
}
