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
// time   : 2020-04-10 9:14
// version: 1.0.0
// desc   : 配置类

package gog

// Config 配置类
type Config struct {
	Formatter Formatter // 格式化
	Writers   []Writer  // 输出器
}

var (
	// 默认配置
	defaultConfig = &Config{
		Formatter: NewNormalColorfulFormatter(),
		Writers:   []Writer{NewConsoleWriter()},
	}
)

// SetFormatter 设置格式化
func (c *Config) SetFormatter(ftr Formatter) {
	c.Formatter = ftr
}

// SetWriter 设置输出器
func (c *Config) SetWriter(wtr ...Writer) {
	c.Writers = append([]Writer{}, wtr...)
}

// AddWriter 添加输出器
func (c *Config) AddWriter(wtr ...Writer) {
	c.Writers = append(c.Writers, wtr...)
}

// ResetWriter 重置输出器
func (c *Config) ResetWriter() {
	c.Writers = make([]Writer, 0)
}
