package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	metriccommon "tcloud_exporter/metriccommon"
	"tcloud_exporter/metrics"
	"tcloud_exporter/utils"
	"time"
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
	mysqlMemUseRage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tcloud",
		Subsystem: "database_mysql",
		Name: "mem",
		Help: "Number of blob storage operations waiting to be processed.",
	},
		[]string{"instance"})
	mysqlDiskUse = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tcloud",
		Subsystem: "database_mysql",
		Name: "disk",
		Help: "Number of blob storage operations waiting to be processed.",
	},
		[]string{"instance"})
	mysqlNetIn = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tcloud",
		Subsystem: "database_mysql",
		Name: "net_in",
		Help: "Number of blob storage operations waiting to be processed.",
	},
		[]string{"instance"})
	mysqlNetOut = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tcloud",
		Subsystem: "database_mysql",
		Name: "net_out",
		Help: "Number of blob storage operations waiting to be processed.",
	},
		[]string{"instance"})
)

func init(){
	prometheus.MustRegister(mysqlCpuUseRage)
	prometheus.MustRegister(mysqlMemUseRage)
	prometheus.MustRegister(mysqlDiskUse)
	prometheus.MustRegister(mysqlNetIn)
	prometheus.MustRegister(mysqlNetOut)
}

func main(){
	InitConfig()
	TENCENTCLOUD_SECRET_ID,TENCENTCLOUD_SECRET_KEY = utils.GetAuthInfo()
	go func(){
		for {
			mysqlmetrics := metrics.GetMysqlMetrics(TENCENTCLOUD_SECRET_ID,TENCENTCLOUD_SECRET_KEY)
			metriccommon.GetGuage("CPUUseRate",mysqlCpuUseRage,mysqlmetrics)
			metriccommon.GetGuage("MemoryUseRate",mysqlMemUseRage,mysqlmetrics)
			metriccommon.GetGuage("VolumeRate",mysqlDiskUse,mysqlmetrics)
			metriccommon.GetGuage("BytesSent",mysqlNetOut,mysqlmetrics)
			metriccommon.GetGuage("BytesReceived",mysqlNetIn,mysqlmetrics)
			time.Sleep(time.Second * 60)

		}
	}()

	flag.Parse()
	http.Handle("/metrics",promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr,nil))

}