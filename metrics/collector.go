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
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
)

// 根据当前配置信息，获取配置里面的数据库项，并根据数据库项获取响应的数据库指标
func GetResourceList(id, key string, resourceconfig *viper.Viper, dataconfig *viper.Viper) {
	objects := dataconfig.AllKeys()
	for _, val := range objects {
		data := dataconfig.GetStringSlice(val)
		fmt.Println(data)
	}
}

func GetResourceMetric(id, key string, resourceconfig *viper.Viper, dataconfig *viper.Viper) {
	client := GetClient(id, key)
	GetMysqlMetrics(client, resourceconfig, dataconfig)
}

func GetClient(id, key string) *monitor.Client {
	cpf := GetCpf()
	// 认证信息
	credential := common.NewCredential(id, key)
	client, _ := monitor.NewClient(credential, regions.Beijing, cpf)
	return client

}
