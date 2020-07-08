package metrics

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"tcloud_exporter/utils"
)

type Data struct{
	Key string
	Value float64
}
type MetricObj struct{
	MetricData map[string][]Data
	//MetricName string
	//Data []Data
}

func GetMetrics(client *monitor.Client,MetricCollector *MetricObj,apinamespace string ,metrictype string){
	// 创建并设置请求参数
	request := monitor.NewGetMonitorDataRequest()
	request.Namespace = common.StringPtr(apinamespace)
	request.MetricName = common.StringPtr(metrictype)
	//设置采集时间
	utils.SetTimeRange(request)
	// instance 设置
	AddInstance(request)
	// 发起请求
	response,err := client.GetMonitorData(request)
	// 异常处理
	if _,ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned : %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	FormatMetrics(response,MetricCollector)
}

func FormatMetrics(response *monitor.GetMonitorDataResponse,MetricCollector *MetricObj){
	//metrics := make(map[string]float64)
	datas := make([]Data,0)
	for _,i := range response.Response.DataPoints {

		instanceid := *i.Dimensions[0].Value
		value := *i.Values[0]
		data := Data{instanceid,value}
		datas = append(datas,data)
	}
	MetricCollector.MetricData[*response.Response.MetricName] =  datas
}

func AddInstance(request *monitor.GetMonitorDataRequest){
	mysqllist := utils.GetMysqlInstance()
	list_instance := []*monitor.Instance{}
	for _,str := range mysqllist {
		list_dimension := []*monitor.Dimension{}
		dimension := &monitor.Dimension{common.StringPtr("InstanceId"),common.StringPtr(str)}
		list_dimension = append(list_dimension,dimension)
		instance := &monitor.Instance{list_dimension}
		list_instance = append(list_instance,instance)

	}
	request.Instances = list_instance
}
