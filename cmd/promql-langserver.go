// Copyright 2019 Tobias Guggenmos
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/slrtbtfs/promql-lsp/langserver"
)

func main() {
	pwd, _ := os.Getwd()
	config, err := langserver.ParseConfigFile("promql-lsp.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading config file:", err.Error(), pwd)
		os.Exit(1)
	}
	ctx, s := langserver.StdioServer(context.Background(), config)
	s.Run(ctx)
}
