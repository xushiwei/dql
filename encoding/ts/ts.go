/*
 * Copyright (c) 2026 The XGo Authors (xgo.dev). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ts

import (
	"strings"

	"github.com/goplus/dql/ts"
)

const (
	XGoPackage = "github.com/goplus/xgo/dql/ts"
)

// Object represents a TypeScript File.
type Object = *ts.File

// New parses TypeScript source code from the given text, returning a TypeScript
// File object. An optional Config can be provided to customize the parsing behavior.
func New(text string, conf ...ts.Config) (f Object, err error) {
	return ts.ParseFile("index.ts", strings.NewReader(text), conf...)
}
