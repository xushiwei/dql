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

package xgo

import (
	"iter"
	"reflect"

	"github.com/microsoft/typescript-go/ast"
	"github.com/goplus/xgo/dql/reflects"
)

const (
	XGoPackage = "github.com/goplus/xgo/dql/reflects"
)

// -----------------------------------------------------------------------------

// Node represents a XGo AST node.
type Node = reflects.Node

// NodeSet represents a set of XGo AST nodes.
type NodeSet struct {
	reflects.NodeSet
}

// NodeSet(seq) casts a NodeSet from a sequence of nodes.
func NodeSet_Cast(seq iter.Seq[Node]) NodeSet {
	return NodeSet{
		NodeSet: reflects.NodeSet{Data: seq},
	}
}

// Root creates a NodeSet containing the provided root node.
func Root(doc Node) NodeSet {
	return NodeSet{
		NodeSet: reflects.Root(doc),
	}
}

// Nodes creates a NodeSet containing the provided nodes.
func Nodes(nodes ...Node) NodeSet {
	return NodeSet{
		NodeSet: reflects.Nodes(nodes...),
	}
}

// New creates a NodeSet from the given *ast.SourceFile.
func New(f *ast.SourceFile) NodeSet {
	return NodeSet{
		NodeSet: reflects.New(reflect.ValueOf(f)),
	}
}

// -----------------------------------------------------------------------------
