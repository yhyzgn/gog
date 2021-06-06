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
// time   : 2020-04-10 9:26
// version: 1.0.0
// desc   : JSON格式化

package gog

import (
	"encoding/json"
	"strconv"
	"strings"
)

// JSONFormatter json 格式化
type JSONFormatter struct {
	Pretty     bool   // 是否美化 json 数据
	TimeLayout string // 日期时间格式
}

// json 日志的一些字段
type jsonLog struct {
	Tag       string      `json:"tag"`
	Timestamp string      `json:"timestamp"`
	Level     string      `json:"level"`
	Func      string      `json:"func"`
	Message   interface{} `json:"message"`
}

// newJSONFormatter 创建 json 格式化对象
func newJSONFormatter(pretty bool) *JSONFormatter {
	return &JSONFormatter{
		Pretty:     pretty,
		TimeLayout: DatePattern,
	}
}

// NewJSONFormatter 创建 json 格式化对象
func NewJSONFormatter() *JSONFormatter {
	return newJSONFormatter(false)
}

// NewJSONPrettyFormatter 创建 json 格式化对象
func NewJSONPrettyFormatter() *JSONFormatter {
	return newJSONFormatter(true)
}

// Format 具体的格式化定义
func (jf *JSONFormatter) Format(level Level, levelName string, info *LogInfo) ([]byte, error) {
	log := jsonLog{
		Tag:       info.Tag,
		Timestamp: info.Time.Format(DatePattern),
		Level:     levelName,
		Func:      info.File + ":" + strconv.Itoa(info.Line) + " (" + info.Func + ")",
		Message:   info.Body,
	}

	var (
		bs  []byte
		err error
	)
	if jf.Pretty {
		bs, err = json.MarshalIndent(log, "", "\t")
	} else {
		bs, err = json.Marshal(log)
	}
	if err != nil {
		return nil, err
	}
	var sb strings.Builder
	sb.Write(bs)
	sb.WriteString("\n")
	return []byte(sb.String()), nil
}
