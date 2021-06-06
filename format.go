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
// time   : 2019-11-04 11:39 下午
// version: 1.0.0
// desc   : 日志格式化类

package gog

import (
	"github.com/yhyzgn/golus"
	"strings"
)

// WithConnectors 在某字符串前拼接 ${length} 个连接符，连接符个数为 length - len(item)
// 为了保持 右对齐
func WithConnectors(item, connector string, length int) string {
	delta := length - len(item)
	if delta <= 0 {
		return item
	}
	var sb strings.Builder
	sb.WriteString(" ")
	for i := 0; i < delta; i++ {
		sb.WriteString(connector)
	}
	sb.WriteString(" ")
	sb.WriteString(item)
	return sb.String()
}

// LevelStylus 给不同 level 的日志加上色彩风格
func LevelStylus(lvl Level) *golus.Stylus {
	stylus := golus.NewStylus()
	switch lvl {
	case TRACE:
		stylus.SetFontColor(golus.FontPurple)
		break
	case DEBUG:
		stylus.SetFontColor(golus.FontBlueGreen)
		break
	case INFO:
		stylus.SetFontColor(golus.FontGreen)
		break
	case WARN:
		stylus.SetFontColor(golus.FontYellow)
		break
	case ERROR:
		stylus.SetFontColor(golus.FontRed)
		break
	case FATAL:
		stylus.SetFontColor(golus.FontRed).SetFontStyle(golus.StyleBold)
		break
	}
	return stylus
}
