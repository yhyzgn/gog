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
// time   : 2019-11-04 8:44 下午
// version: 1.0.0
// desc   : 对外日志工具

package gog

import (
	"sync"
)

var (
	once sync.Once
	gog  *Gog // 默认的日志对象
)

func init() {
	once.Do(func() {
		// 每多封装一层，就需要多跳过一层调用栈
		gog = NewGog(ALL, 1)
		// 默认只显示文件名，不显示完整路径
		gog.ShortFile(true)
	})
}

// SetGog 设置日志对象
func SetGog(g *Gog) {
	gog = g
}

// GetGog 获取日志对象
func GetGog() *Gog {
	return gog
}

// SetConfig 设置配置
func SetConfig(cfg *Config) {
	gog.SetConfig(cfg)
}

// SetFormatter 设置格式化模式
func SetFormatter(ftr Formatter) {
	gog.SetFormatter(ftr)
}

// SetWriter 设置输出器
func SetWriter(wtr ...Writer) {
	gog.SetWriter(wtr...)
}

// AddWriter 添加输出器
func AddWriter(wtr ...Writer) {
	gog.AddWriter(wtr...)
}

// ResetWriters 重置输出器
func ResetWriters() {
	gog.ResetWriters()
}

// CallSkip 设置需要跳过的调用栈
//
// 如果要将 'gog.Info(...)' 再封装一层，就需要设置 skip 值，日志输出时才能打印出具体的调用地方
// skip 最小值为 1
func CallSkip(skip int) {
	gog.CallSkip(skip)
}

// SetLevel 设置日志打印的最低优先级
func SetLevel(lvl Level) {
	gog.Level(lvl)
}

// ShortFile 是否只显示文件名
func ShortFile(short bool) {
	gog.ShortFile(short)
}

// Async 是否启用异步
//
// 默认关闭
func Async(async bool) {
	gog.Async(async)
}

// Trace 追踪打印
func Trace(value ...interface{}) {
	gog.Trace(value...)
}

// TraceTag 追踪打印
func TraceTag(tag string, value ...interface{}) {
	gog.TraceTag(tag, value...)
}

// TraceF 追踪打印
func TraceF(format string, args ...interface{}) {
	gog.TraceF(format, args...)
}

// TraceTagF 追踪打印
func TraceTagF(tag string, format string, args ...interface{}) {
	gog.TraceTagF(tag, format, args...)
}

// Debug 调试打印
func Debug(value ...interface{}) {
	gog.Debug(value...)
}

// DebugTag 调试打印
func DebugTag(tag string, value ...interface{}) {
	gog.DebugTag(tag, value...)
}

// DebugF 调试打印
func DebugF(format string, args ...interface{}) {
	gog.DebugF(format, args...)
}

// DebugTagF 调试打印
func DebugTagF(tag string, format string, args ...interface{}) {
	gog.DebugTagF(tag, format, args...)
}

// Info 普通信息打印
func Info(value ...interface{}) {
	gog.Info(value...)
}

// InfoTag 普通信息打印
func InfoTag(tag string, value ...interface{}) {
	gog.InfoTag(tag, value...)
}

// InfoF 普通信息打印
func InfoF(format string, args ...interface{}) {
	gog.InfoF(format, args...)
}

// InfoTagF 普通信息打印
func InfoTagF(tag string, format string, args ...interface{}) {
	gog.InfoTagF(tag, format, args...)
}

// Warn 警告打印
func Warn(value ...interface{}) {
	gog.Warn(value...)
}

// WarnTag 警告打印
func WarnTag(tag string, value ...interface{}) {
	gog.WarnTag(tag, value...)
}

// WarnF 警告打印
func WarnF(format string, args ...interface{}) {
	gog.WarnF(format, args...)
}

// WarnTagF 警告打印
func WarnTagF(tag string, format string, args ...interface{}) {
	gog.WarnTagF(tag, format, args...)
}

// Error 错误打印
func Error(value ...interface{}) {
	gog.Error(value...)
}

// ErrorTag 错误打印
func ErrorTag(tag string, value ...interface{}) {
	gog.ErrorTag(tag, value...)
}

// ErrorF 错误打印
func ErrorF(format string, args ...interface{}) {
	gog.ErrorF(format, args...)
}

// ErrorTagF 错误打印
func ErrorTagF(tag string, format string, args ...interface{}) {
	gog.ErrorTagF(tag, format, args...)
}

// Fatal 错误打印，并结束进程
func Fatal(value ...interface{}) {
	gog.Fatal(value...)
}

// FatalTag 错误打印，并结束进程
func FatalTag(tag string, value ...interface{}) {
	gog.FatalTag(tag, value...)
}

// FatalF 错误打印，并结束进程
func FatalF(format string, args ...interface{}) {
	gog.FatalF(format, args...)
}

// FatalTagF 错误打印，并结束进程
func FatalTagF(tag string, format string, args ...interface{}) {
	gog.FatalTagF(tag, format, args...)
}

// Log 适配器
func Log(level Level, args ...interface{}) {
	switch level {
	case TRACE:
		gog.Trace(args...)
		break
	case DEBUG:
		gog.Debug(args...)
		break
	case INFO:
		gog.Info(args...)
		break
	case WARN:
		gog.Warn(args...)
		break
	case ERROR:
		gog.Error(args...)
		break
	case FATAL:
		gog.Fatal(args...)
		break
	default:
		gog.Info(args...)
	}
}

// LogF 适配器
func LogF(level Level, format string, args ...interface{}) {
	switch level {
	case TRACE:
		gog.TraceF(format, args...)
		break
	case DEBUG:
		gog.DebugF(format, args...)
		break
	case INFO:
		gog.InfoF(format, args...)
		break
	case WARN:
		gog.WarnF(format, args...)
		break
	case ERROR:
		gog.ErrorF(format, args...)
		break
	case FATAL:
		gog.FatalF(format, args...)
		break
	default:
		gog.InfoF(format, args...)
	}
}
