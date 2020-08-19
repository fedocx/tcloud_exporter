package utils

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"time"
)

func SetTimeRange(request *monitor.GetMonitorDataRequest){
	//request.StartTime = collector.StringPtr(time.Now().Add(time.Duration(-1)*time.Minute).Format(time.RFC3339))
	//request.EndTime = common.StringPtr(time.Now().Format(time.RFC3339))
	request.StartTime = common.StringPtr(time.Now().Format(time.RFC3339))
	m,_ := time.ParseDuration("10m")
	request.EndTime = common.StringPtr(time.Now().Add(m).Format(time.RFC3339))
}