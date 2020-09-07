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
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"log"
	"tcloud_exporter/utils"
)

// struct for collector
type Data struct {
	Key   string
	Value float64
}

type Product struct{
	Metrics map[string][]*Data
}


// struct for collector
type MetricObj struct {
	Products  map[string][]*Product
}


func GetMetrics(client *monitor.Client, MetricCollector *MetricObj, value_temp MetricChannel,lock chan int) {
	<- lock
	//参数初始化
	apinamespace := value_temp.Apinamespace
	metrictype := value_temp.MetricType
	objecttype := value_temp.Type
	//instancename := value_temp.InstanceName
	config := value_temp.Config

	// log
	log.Print("开始采集:", objecttype, "    ",metrictype )
	// 创建并设置请求参数
	request := monitor.NewGetMonitorDataRequest()
	request.Namespace = common.StringPtr(apinamespace)
	request.MetricName = common.StringPtr(metrictype)
	request.Period = common.Uint64Ptr(300)
	//设置采集时间
	utils.SetTimeRange(request)
	// instance 设置

	// print request delete
	//fmt.Println(apinamespace,metrictype,instancelist)

	//AddInstance(request, instancelist)
	AddInstance(request,config)
	// 发起请求
	response, err := client.GetMonitorData(request)
	// 异常处理
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Println(*request.Namespace,*request.MetricName,*request.Instances[0].Dimensions[0])
		fmt.Printf("An API error has returned : %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	FormatMetrics(objecttype,response, MetricCollector)
	lock <- 1
}

// 客户端配置对象
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

func FormatMetrics(productname string,response *monitor.GetMonitorDataResponse, MetricCollector *MetricObj) {
	//metrics := make(map[string]float64)
	datas := make([]*Data, 0)
	for _, i := range response.Response.DataPoints {

		instanceid := *i.Dimensions[0].Value
		var value float64
		if len(i.Values) > 0 {
			value = *i.Values[len(i.Values) -1]
		} else {
			continue
		}
		data := Data{instanceid, value}
		datas = append(datas, &data)
	}
	//metrics := new(Product)
	metrics := make(map[string][]*Data)
	Metrics := new(Product)
	Metrics.Metrics = metrics


	Metrics.Metrics[*response.Response.MetricName] = datas
	log.Print("采集到数据:", metrics, ":",datas)
	//productname = NamespaceToNameMap(productname)
	MetricCollector.Products[productname] = append(MetricCollector.Products[productname],Metrics)
	//MetricCollector.Products[productname].[*response.Response.MetricName] = datas
	//fmt.Println(MetricCollector.Products)
}

func AddInstance(request *monitor.GetMonitorDataRequest, instancelist []map[string]string) {
	//mysqllist := utils.GetMysqlInstance(resourceconfig)
	list_instance := []*monitor.Instance{}
	for _, str := range instancelist {
		list_dimension := []*monitor.Dimension{}
		var dimension *monitor.Dimension
		for key,val := range str{
			dimension = &monitor.Dimension{common.StringPtr(key), common.StringPtr(val)}
			list_dimension = append(list_dimension, dimension)
		}
		instance := &monitor.Instance{list_dimension}
		list_instance = append(list_instance, instance)

	}
	request.Instances = list_instance
}


// 通过id，key生成客户端
func GetClient(id, key string) *monitor.Client {
	cpf := GetCpf()
	// 认证信息
	credential := common.NewCredential(id, key)
	client, _ := monitor.NewClient(credential, regions.Beijing, cpf)
	return client

}
