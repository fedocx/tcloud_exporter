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

type Mysql struct {
}

//func GetMysqlMetrics(client *monitor.Client,MetricCollector *MetricObj, resourceconfig *viper.Viper, dataconfig *viper.Viper) *MetricObj {
//
//
//	//mysql register
//	mysqlmetrics := utils.GetMysqlMetrics(dataconfig)
//
//	// 获取指标
//	for _, val := range mysqlmetrics {
//		fmt.Println(val)
//		GetMysqlMetric(client, MetricCollector, val, resourceconfig)
//	}
//	return MetricCollector
//}

//func GetMysqlMetric(client *monitor.Client, MetricCollector *MetricObj, metrictype string, resourceconfig *viper.Viper) {
//	GetMetrics(client, MetricCollector, "QCE/CDB", metrictype, resourceconfig)
//}

func GetMysqlCode()string{
	return "QCE/CDB"
}

//func SendMysqlMetric(client *monitor.Client, MetricCollector *MetricObj, metrictype string, resourceconfig *viper.Viper){
//
//}