package metriccommon

import (
	"github.com/prometheus/client_golang/prometheus"
	"tcloud_exporter/metrics"
)

func GetGuage(metricname string,gaugevec *prometheus.GaugeVec,metrics *metrics.MetricObj){
	for _,val := range metrics.MetricData[metricname] {
		i,_ := gaugevec.GetMetricWithLabelValues(val.Key)
		i.Set(val.Value)
	}

}