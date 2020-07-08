package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"tcloud_exporter/metrics"
	"tcloud_exporter/utils"
	"time"
	"log"
)

var TENCENTCLOUD_SECRET_ID,TENCENTCLOUD_SECRET_KEY string
func InitConfig(){
	workDir,_ := os.Getwd()
	viper.SetConfigName("tencent")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

var addr = flag.String("listen-addr",":8081","the port to listen on for HTTP requests")

var (
	mysqlCpuUseRage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tcloud",
		Subsystem: "database_mysql",
		Name: "cpu",
		Help: "Number of blob storage operations waiting to be processed.",
	},
	[]string{"instance"})
	//mysqlMemUseRage = prometheus.NewHistogram(prometheus.HistogramOpts{
	//	Namespace: "tcloud",
	//	Subsystem: "blob_storage",
	//	Name: "memory",
	//	Help: "Number of blob storage operations waiting to be processed.",
	//	Buckets: prometheus.ExponentialBuckets(1,5,5),
	//})
)

func init(){
	prometheus.MustRegister(mysqlCpuUseRage)
}

func main(){
	InitConfig()
	TENCENTCLOUD_SECRET_ID,TENCENTCLOUD_SECRET_KEY = utils.GetAuthInfo()
	go func(){
		for {
			mysqlmetrics := *metrics.GetMysqlMetrics(TENCENTCLOUD_SECRET_ID,TENCENTCLOUD_SECRET_KEY)
			for _,val := range mysqlmetrics.MetricData["CPUUseRate"]{
				fmt.Println(val.Key,val.Value)
				//mysqlCpuUseRage.With(prometheus.Labels{"instance":val.Key})
				i,_ := mysqlCpuUseRage.GetMetricWithLabelValues(val.Key)
				i.Set(val.Value)
			}
			time.Sleep(time.Second * 60)

		}
	}()

	flag.Parse()
	http.Handle("/metrics",promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr,nil))

}