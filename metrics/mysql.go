package metrics

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
)

type Mysql struct {

}


func GetMysqlMetrics(id,key string)(*MetricObj){
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
	return MetricCollector


}

func GetMysqlMetric(client *monitor.Client,MetricCollector *MetricObj,metrictype string){
	GetMetrics(client,MetricCollector,"QCE/CDB",metrictype)
}