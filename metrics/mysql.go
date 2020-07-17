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
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
)

type Mysql struct {

}


func GetMysqlMetrics(id,key string)(*MetricObj){

	cpf := GetCpf()
    // 认证信息
	credential := common.NewCredential(id,key)
	client,_ := monitor.NewClient(credential,regions.Beijing,cpf)

	MetricCollector := new(MetricObj)
	Metricdata := make(map[string][]Data)
	MetricCollector.MetricData = Metricdata

	// 获取指标
	GetMysqlMetric(client,MetricCollector,"CPUUseRate")
	GetMysqlMetric(client,MetricCollector,"MemoryUseRate")
	GetMysqlMetric(client,MetricCollector,"BytesSent")
	GetMysqlMetric(client,MetricCollector,"BytesReceived")
	GetMysqlMetric(client,MetricCollector,"VolumeRate")
	return MetricCollector
}

func GetMysqlMetric(client *monitor.Client,MetricCollector *MetricObj,metrictype string){
	GetMetrics(client,MetricCollector,"QCE/CDB",metrictype)
}

