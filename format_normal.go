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
// time   : 2020-04-10 9:19
// version: 1.0.0
// desc   : 普通日志格式化模式

package gog

import (
	"github.com/yhyzgn/gog/util"
	"strconv"
	"strings"
)

// NormalFormatter 控制台格式化
type NormalFormatter struct {
	Colorful   bool   // 是否支持五颜六色
	TimeLayout string // 日期时间格式
}

// newNormalFormatter 创建控制台格式化对象
func newNormalFormatter(colorful bool) *NormalFormatter {
	return &NormalFormatter{
		Colorful:   colorful,
		TimeLayout: DatePattern,
	}
}

// NewNormalFormatter 创建控制台格式化对象
func NewNormalFormatter() *NormalFormatter {
	return newNormalFormatter(false)
}

// NewNormalColorfulFormatter 创建控制台格式化对象
func NewNormalColorfulFormatter() *NormalFormatter {
	return newNormalFormatter(true)
}

// Format 具体的格式化定义
func (cf *NormalFormatter) Format(level Level, levelName string, info *LogInfo) ([]byte, error) {
	var sb strings.Builder
	sb.WriteString(info.Time.Format(DatePattern))
	sb.WriteString(WithConnectors(levelName, "-", 8))
	sb.WriteString(WithConnectors(info.File, "-", util.If(info.ShortFile, FileLengthRel, FileLengthAbs).(int)))
	sb.WriteString(":" + util.FillSuffix(strconv.Itoa(info.Line), " ", 4))
	sb.WriteString("(" + info.Func + ")")
	if info.Tag != "" {
		sb.WriteString("[" + info.Tag + "]")
	}
	sb.WriteString(info.Body)

	res := sb.String()
	if cf.Colorful {
		res = LevelStylus(level).Apply(res)
	}
	return []byte(res + "\n"), nil
}
