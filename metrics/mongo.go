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

import "github.com/spf13/viper"

type Mongodb struct {
}

type Mongodb_instance struct{
	Target string
}

func (t *Mongodb) GetCode() string {
	return "QCE/CMONGO"
}

func(t *Mongodb) GetMetrics(dataconfig *viper.Viper) []string{
	return dataconfig.GetStringSlice("mongodb")
}

type Mongodb_cluster struct {
}


func (t *Mongodb_cluster) GetCode() string {
	return "QCE/CMONGO"
}

func(t *Mongodb_cluster) GetMetrics(dataconfig *viper.Viper) []string{
	return dataconfig.GetStringSlice("mongodb_cluster")
}
type Mongodb_replication struct {
}


func (t *Mongodb_replication) GetCode() string {
	return "QCE/CMONGO"
}

func(t *Mongodb_replication) GetMetrics(dataconfig *viper.Viper) []string{
	return dataconfig.GetStringSlice("mongodb_replication")
}
