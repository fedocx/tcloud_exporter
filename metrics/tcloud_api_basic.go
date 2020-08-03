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
// file for set tencent api parameters
package metrics

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"tcloud_exporter/utils"
)

// struct for collector
type Data struct {
	Key   string
	Value float64
}

// struct for collector
type MetricObj struct {
	MetricData map[string][]Data
	//MetricName string
	//Data []Data
}

func GetMetrics(client *monitor.Client, MetricCollector *MetricObj, apinamespace string, metrictype string, resourceconfig *viper.Viper) {
	// 创建并设置请求参数
	request := monitor.NewGetMonitorDataRequest()
	request.Namespace = common.StringPtr(apinamespace)
	request.MetricName = common.StringPtr(metrictype)
	//设置采集时间
	utils.SetTimeRange(request)
	// instance 设置
	AddInstance(request, resourceconfig)
	// 发起请求
	response, err := client.GetMonitorData(request)
	// 异常处理
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned : %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	FormatMetrics(response, MetricCollector)
}

func GetCpf() *profile.ClientProfile {
	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用POST方法。
	// 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
	cpf.HttpProfile.ReqMethod = "POST"
	// SDK有默认的超时时间，非必要请不要进行调整。
	// 如有需要请在代码中查阅以获取最新的默认值。
	cpf.HttpProfile.ReqTimeout = 10
	// SDK会自动指定域名。通常是不需要特地指定域名的，但是如果你访问的是金融区的服务，
	// 则必须手动指定域名，例如云服务器的上海金融区域名： cvm.ap-shanghai-fsi.tencentcloudapi.com
	cpf.HttpProfile.Endpoint = "monitor.ap-beijing.tencentcloudapi.com"
	return cpf

}

func FormatMetrics(response *monitor.GetMonitorDataResponse, MetricCollector *MetricObj) {
	//metrics := make(map[string]float64)
	datas := make([]Data, 0)
	for _, i := range response.Response.DataPoints {

		instanceid := *i.Dimensions[0].Value
		var value float64
		if len(i.Values) > 0 {
			value = *i.Values[0]
		} else {
			continue
		}
		data := Data{instanceid, value}
		datas = append(datas, data)
	}
	MetricCollector.MetricData[*response.Response.MetricName] = datas
}

func AddInstance(request *monitor.GetMonitorDataRequest, resourceconfig *viper.Viper) {
	mysqllist := utils.GetMysqlInstance(resourceconfig)
	list_instance := []*monitor.Instance{}
	for _, str := range mysqllist {
		list_dimension := []*monitor.Dimension{}
		dimension := &monitor.Dimension{common.StringPtr("InstanceId"), common.StringPtr(str)}
		list_dimension = append(list_dimension, dimension)
		instance := &monitor.Instance{list_dimension}
		list_instance = append(list_instance, instance)

	}
	request.Instances = list_instance
}
