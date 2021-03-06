/*
 * Copyright 2022 steven
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *    http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package simBench

import "github.com/spf13/pflag"

var (
	concurrency int  // 并发请求数量 ，比如 -c 100 ，代表每秒创建100个链接
	number      int // 总的请求数量 ，比如 -n 10000,代表总共建立链接10000个
	tagNum      int // 随机生成tag的数量
	keepTime    int // 总的在线时长，比如 -k 100 ,代表在线时间为100秒，100秒后就会释放
	host        string // 设置对应的Url ，比如 -h 127.0.0.1:8080,目前展示不能支持配置链接token，后续会加上
)

// 整体流程： connection  -c -
func init() {
	pflag.IntVarP(&concurrency, "concurrency", "c", 100, "并发请求数量 ，比如 -c 100 ，代表每秒创建100个链接")
	pflag.IntVarP(&number, "number", "n", 1000, "总的请求数量 ，比如 -n 10000,代表总共建立链接10000个")
	pflag.IntVarP(&tagNum, "tags", "t", 3, "每个连接的默认的tag数量 ")
	pflag.IntVarP(&keepTime, "keepTime", "k", 0, "总的在线时长，比如 -k 100s ,代表在线时间为100秒，100秒后就会释放，为0不释放")
	pflag.StringVarP(&host, "host", "h", "ws://127.0.0.1:8081/conn", "设置对应的Url")
	pflag.Parse()
}

type Config struct {
	Concurrency int
	Number      int
	KeepTime    int
	Host        string
	TagNum 		int
}

// InitConfig 实例化
func InitConfig()*Config {
	return &Config{
		Concurrency: concurrency,
		Number:      number,
		KeepTime:    keepTime,
		Host:        host,
		TagNum:      tagNum,
	}
}
