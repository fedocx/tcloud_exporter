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

type Kafka_topic struct {
}

func (t *Kafka_topic) GetCode() string {
	return "QCE/CKAFKA"
}
func(t *Kafka_topic) GetMetrics(dataconfig *viper.Viper) []string{
	return dataconfig.GetStringSlice("kafka_topic")
}



type Kafka_partition struct {
}

func (t *Kafka_partition) GetCode() string {
	return "QCE/CKAFKA"
}

func(t *Kafka_partition) GetMetrics(dataconfig *viper.Viper) []string{
	return dataconfig.GetStringSlice("kafka_partition")

}

