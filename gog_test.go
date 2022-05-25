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
	"fmt"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	_ = os.Setenv("ENV", "dev")

	env := os.Getenv("ENV")
	if env == "prod" {
		SetFormatter(NewJSONFormatter())
		// 开发环境
		AddWriter(&testWriter{})
	}

	SetLevel(INFO)

	Async(true)

	Trace("Trace", 1, 3, 345, true)
	Debug("Debug")
	Info("Info")
	Warn("Warn")
	Error("Error")

	time.Sleep(3 * time.Second)

	TraceTag("TagTrace", "Trace", 1, 3, 345, true)
	DebugTag("TagDebug", "Debug")
	InfoTag("TagInfo", "Info")
	WarnTag("TagWarn", "Warn")
	ErrorTagF("TagErrorF", "Error {}", "err...")

	//Fatal("Fatal")

	test()

	time.Sleep(6 * time.Second)
}

func test() {
	TraceF("Hello, {} !", "gog")
}

type testWriter struct {
	ConsoleWriter
}

func (tw *testWriter) Write(info *LogInfo, data []byte) (n int, err error) {
	if info.Level >= INFO {
		fmt.Print("test writer -- ", string(data))
	}
	return
}
