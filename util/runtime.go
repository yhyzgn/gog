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
// time   : 2019-11-04 11:47 下午
// version: 1.0.0
// desc   : runtime 工具

package util

import (
	"path/filepath"
	"runtime"
	"strings"
)

// FileLineNumber 获取调用源的   文件 方法 行号   信息
func FileLineNumber(callSkip int, shortFile bool) (string, string, int, bool) {
	fnPtr, file, line, ok := runtime.Caller(callSkip)
	if ok {
		funcName := runtime.FuncForPC(fnPtr).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")
		if shortFile {
			file = filepath.Base(file)
		}
		return file, funcName, line, true
	}
	return "", "", 0, false
}
