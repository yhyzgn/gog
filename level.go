// Copyright 2019 yhyzgn gog
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-11-05 1:03 上午
// version: 1.0.0
// desc   : 日志级别定义

package gog

import "strings"

// Level 日志级别类型
type Level int

// 一些日志级别，优先级递增
const (
	ALL   Level = iota // 最低等级，打开所有日志级别
	TRACE              // 追踪级别
	DEBUG              // 调试级别
	INFO               // 一般级别
	WARN               // 警告级别
	ERROR              // 错误级别，打印错误，程序继续运行
	FATAL              // 严重错误，将导致程序退出
	OFF                // 关闭所有日志
)

// 日志级别名称
var levelNames = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

// GetLevelName 获取日志级别名称
func GetLevelName(level Level) string {
	if level > 0 && int(level) <= len(levelNames) {
		return levelNames[level-1]
	}
	return "UNKNOWN"
}

// ParseLevel 根据 name 解析 level
func ParseLevel(levelName string) Level {
	switch strings.ToUpper(levelName) {
	case "TRACE":
		return TRACE
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	case "OFF":
		return OFF
	}
	return ALL
}
