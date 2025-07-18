// Copyright 2022 The etcd Authors
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

//go:build tools
// +build tools

// This file implements that pattern:
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// for raft. Thanks to this file 'go mod tidy' does not remove dependencies.

package tools

import (
	_ "github.com/alexkohler/nakedret"
	_ "github.com/chzchzchz/goword"
	_ "github.com/coreos/license-bill-of-materials"
	_ "github.com/google/addlicense"
	_ "github.com/gordonklaus/ineffassign"
	_ "github.com/gyuho/gocovmerge"
	_ "github.com/hexfusion/schwag"
	_ "github.com/mdempsky/unconvert"
	_ "github.com/mgechev/revive"
	_ "github.com/mikefarah/yq/v4"
	_ "gotest.tools/gotestsum"
	_ "gotest.tools/v3"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "mvdan.cc/unparam"

	_ "go.etcd.io/gofail"
	_ "go.etcd.io/protodoc"
)
