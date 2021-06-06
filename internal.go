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
// time   : 2020-04-10 9:12
// version: 1.0.0
// desc   : 内置定义

package gog

import (
	"io"
	"time"
)

const (
	// DatePattern 日期格式化格式
	DatePattern = "2006-01-02 15:04:05.999"
	// FileLengthRel 文件相对路径格式化长度
	FileLengthRel = 26
	// FileLengthAbs 文件绝对路径格式化长度
	FileLengthAbs = 64
)

// Formatter 格式化接口
type Formatter interface {
	Format(level Level, levelName string, info *LogInfo) ([]byte, error)
}

// Writer 日志输出器
type Writer interface {
	io.Closer
	Write(info *LogInfo, data []byte) (n int, err error)
}

// LogInfo 日志数据
type LogInfo struct {
	Tag       string    // 标签
	Time      time.Time // 日志产生时间
	Level     Level     // 日志等级
	Body      string    // 日志详情
	File      string    // 发生地文件
	Func      string    // 发生地函数
	Line      int       // 发生地行号
	ShortFile bool      // 是否为短文件名
}
