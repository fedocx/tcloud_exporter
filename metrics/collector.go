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
	"tcloud_exporter/utils"
)

// 根据当前配置信息，获取配置里面的数据库项，并根据数据库项获取响应的数据库指标
func GetDatabaseMetrics(id,key string,resourceconfig *viper.Viper, dataconfig *viper.Viper){
	objects := dataconfig.AllKeys()
	fmt.Println(objects)
	GetMysqlMetrics(id,key,dastaconfig)

}
