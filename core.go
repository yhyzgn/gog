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
// time   : 2019-11-04 11:18 下午
// version: 1.0.0
// desc   : 日志核心处理器

package gog

import (
	"fmt"
	"github.com/yhyzgn/gog/util"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	queueSize = 100
)

// Gog 日志处理器
type Gog struct {
	mu        sync.Mutex    // 同步锁
	config    *Config       // 配置信息
	callSkip  int           // 定位打印日志的文件、方法以及行号，需要跳过中间调用栈，直接定位到调用源头
	level     Level         // 日志输出级别，只有 >= 该值的级别才会输出
	shortFile bool          // 日志输出时，如果 shortFile=true 则只输出日志源的文件名，否则将输出完整路径
	async     bool          // 是否启用异步
	queue     chan *LogInfo // 异步队列
}

// NewGog 创建新的日志处理器
func NewGog(level Level, callSkip int) *Gog {
	gog := &Gog{
		config:   defaultConfig,
		callSkip: callSkip,
		level:    level,
		queue:    make(chan *LogInfo, queueSize),
	}
	// 开启异步输出
	go gog.startAsyncOut()
	return gog
}

// SetConfig 设置配置
func (g *Gog) SetConfig(cfg *Config) *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config = cfg
	return g
}

// SetFormatter 设置格式化模式
func (g *Gog) SetFormatter(ftr Formatter) *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.SetFormatter(ftr)
	return g
}

// SetWriter 设置输出器
func (g *Gog) SetWriter(wtr Writer) *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.SetWriter(wtr)
	return g
}

// AddWriter 添加输出器
func (g *Gog) AddWriter(wtr Writer) *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.AddWriter(wtr)
	return g
}

// ResetWriters 重置输出器
func (g *Gog) ResetWriters() *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.ResetWriter()
	return g
}

// CallSkip 设置需要跳过的调用栈
//
// 如果要将 'gog.Info(...)' 再封装一层，就需要设置 skip 值，日志输出时才能打印出具体的调用地方
// skip 最小值为 1
func (g *Gog) CallSkip(skip int) *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.callSkip = skip
	return g
}

// Level 设置日志打印的最低优先级
func (g *Gog) Level(level Level) *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.level = level
	return g
}

// ShortFile 是否只显示文件名
func (g *Gog) ShortFile(short bool) *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.shortFile = short
	return g
}

// Async 是否启用异步
//
// 默认关闭
func (g *Gog) Async(async bool) *Gog {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.async = async
	return g
}

// Trace 追踪打印
func (g *Gog) Trace(body ...interface{}) {
	g.Write("", TRACE, body...)
}

// TraceTag 追踪打印
func (g *Gog) TraceTag(tag string, body ...interface{}) {
	g.Write(tag, TRACE, body...)
}

// TraceF 追踪打印
func (g *Gog) TraceF(format string, args ...interface{}) {
	g.Write("", TRACE, resolveFormat(format, args...))
}

// TraceTagF 追踪打印
func (g *Gog) TraceTagF(tag string, format string, args ...interface{}) {
	g.Write(tag, TRACE, resolveFormat(format, args...))
}

// Debug 调试打印
func (g *Gog) Debug(body ...interface{}) {
	g.Write("", DEBUG, body...)
}

// DebugTag 调试打印
func (g *Gog) DebugTag(tag string, body ...interface{}) {
	g.Write(tag, DEBUG, body...)
}

// DebugF 调试打印
func (g *Gog) DebugF(format string, args ...interface{}) {
	g.Write("", DEBUG, resolveFormat(format, args...))
}

// DebugTagF 调试打印
func (g *Gog) DebugTagF(tag string, format string, args ...interface{}) {
	g.Write(tag, DEBUG, resolveFormat(format, args...))
}

// Info 普通信息打印
func (g *Gog) Info(body ...interface{}) {
	g.Write("", INFO, body...)
}

// InfoTag 普通信息打印
func (g *Gog) InfoTag(tag string, body ...interface{}) {
	g.Write(tag, INFO, body...)
}

// InfoF 普通信息打印
func (g *Gog) InfoF(format string, args ...interface{}) {
	g.Write("", INFO, resolveFormat(format, args...))
}

// InfoTagF 普通信息打印
func (g *Gog) InfoTagF(tag string, format string, args ...interface{}) {
	g.Write(tag, INFO, resolveFormat(format, args...))
}

// Warn 警告打印
func (g *Gog) Warn(body ...interface{}) {
	g.Write("", WARN, body...)
}

// WarnTag 警告打印
func (g *Gog) WarnTag(tag string, body ...interface{}) {
	g.Write(tag, WARN, body...)
}

// WarnF 警告打印
func (g *Gog) WarnF(format string, args ...interface{}) {
	g.Write("", WARN, resolveFormat(format, args...))
}

// WarnTagF 警告打印
func (g *Gog) WarnTagF(tag string, format string, args ...interface{}) {
	g.Write(tag, WARN, resolveFormat(format, args...))
}

// Error 错误打印
func (g *Gog) Error(body ...interface{}) {
	g.Write("", ERROR, body...)
}

// ErrorTag 错误打印
func (g *Gog) ErrorTag(tag string, body ...interface{}) {
	g.Write(tag, ERROR, body...)
}

// ErrorF 错误打印
func (g *Gog) ErrorF(format string, args ...interface{}) {
	g.Write("", ERROR, resolveFormat(format, args...))
}

// ErrorTagF 错误打印
func (g *Gog) ErrorTagF(tag string, format string, args ...interface{}) {
	g.Write(tag, ERROR, resolveFormat(format, args...))
}

// Fatal 错误打印，并结束进程
func (g *Gog) Fatal(body ...interface{}) {
	defer func() {
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()
	g.Write("", FATAL, body...)
}

// FatalTag 错误打印，并结束进程
func (g *Gog) FatalTag(tag string, body ...interface{}) {
	defer func() {
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()
	g.Write(tag, FATAL, body...)
}

// FatalF 错误打印，并结束进程
func (g *Gog) FatalF(format string, args ...interface{}) {
	defer func() {
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()
	g.Write("", FATAL, resolveFormat(format, args...))
}

// FatalTagF 错误打印，并结束进程
func (g *Gog) FatalTagF(tag string, format string, args ...interface{}) {
	defer func() {
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()
	g.Write(tag, FATAL, resolveFormat(format, args...))
}

// Write 输出操作
func (g *Gog) Write(tag string, lvl Level, body ...interface{}) {
	if lvl == OFF || lvl < g.level || body == nil || len(body) == 0 {
		return
	}

	bodyFormat := ""
	for range body {
		bodyFormat += "{}"
	}
	info := &LogInfo{
		Tag:       tag,
		Time:      time.Now(),
		Level:     lvl,
		Body:      resolveFormat(bodyFormat, body...),
		ShortFile: g.shortFile,
	}

	// 需要跳过至少3层调用栈
	file, funcName, line, ok := util.FileLineNumber(g.callSkip+3, g.shortFile)
	if ok {
		info.File = file
		info.Func = funcName
		info.Line = line
	}

	if g.async {
		// 添加到异步队列
		select {
		case g.queue <- info:
		default:
			g.out(info)
		}
	} else {
		// 同步输出
		g.out(info)
	}
}

func (g *Gog) out(info *LogInfo) {
	// 设置默认的日志输出器，打印到控制台
	if g.config == nil {
		g.config = defaultConfig
	}

	if len(g.config.Writers) == 0 {
		g.SetWriter(NewConsoleWriter())
	}
	for _, w := range g.config.Writers {
		if g.config.Formatter != nil {
			// 每个输出器自定义输出格式
			data, err := g.config.Formatter.Format(info.Level, GetLevelName(info.Level), info)
			if err != nil {
				log.Fatal(err)
				return
			}
			if _, err := w.Write(info, data); err != nil {
				log.Println(err)
			}
			continue
		}
		log.Println(info)
	}
}

func (g *Gog) startAsyncOut() {
	for {
		g.out(<-g.queue)
	}
}

func resolveFormat(format string, args ...interface{}) string {
	format = strings.ReplaceAll(format, "{}", "%v")
	return fmt.Sprintf(format, args...)
}
