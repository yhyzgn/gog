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
// time   : 2020-04-10 9:28
// version: 1.0.0
// desc   : 控制台输出

package gog

import (
	"io"
	"os"
)

// ConsoleWriter 控制台输出器
type ConsoleWriter struct {
	out io.WriteCloser
}

// NewConsoleWriter 创建控制台输出器对象
func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{
		out: os.Stdout,
	}
}

// Write 输出日志
func (cw *ConsoleWriter) Write(info *LogInfo, data []byte) (n int, err error) {
	return cw.out.Write(data)
}

// 关闭输出流
func (cw *ConsoleWriter) Close() error {
	return cw.out.Close()
}
