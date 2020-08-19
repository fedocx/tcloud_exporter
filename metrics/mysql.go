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
	"fmt"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"reflect"
)

type Mysql struct {
}

type Mysql_instance struct{
	InstanceId string

}

func (t *Mysql) GetCode() string {
	return "QCE/CDB"
}

//func (t *Mysql)GetInstancename()string{
//	return "InstanceId"
//}
//func (t *Mysql) GetInstanceList(config *Config) []Mysql_instance{
//	return config.Mysql
//}
func (t *Mysql) AddInstance(request  *monitor.GetMonitorDataRequest, config *Config){
	list_instance := []*monitor.Instance{}
	t.Rangeinstance(config)
	//for _, str := range config.Kafka {
	//	list_dimension := []*monitor.Dimension{}
	//	for key,val := range str{
	//		dimension := &monitor.Dimension{common.StringPtr(key), common.StringPtr(val)}
	//		list_dimension = append(list_dimension, dimension)
	//	}
	//	instance := &monitor.Instance{list_dimension}
	//	list_instance = append(list_instance, instance)
	//
	//}
	request.Instances = list_instance
}

func (t *Mysql)Rangeinstance(config *Config){
	mysql := config.Mysql
	typ := reflect.TypeOf(mysql)
	val := reflect.ValueOf(mysqlt )
	num := val.NumField()
	for i:=0; i < num; i++{
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != ""{
			fmt.Println(i,tagVal,val.Field(i))
		}
	}

}
