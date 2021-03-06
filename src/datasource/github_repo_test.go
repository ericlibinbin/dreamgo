// Copyright 2017 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris	polaris@studygolang.com

package datasource_test

import (
	"datasource"
	"global"
	"os"
	"strings"

	"testing"
)

func setup() {
	cwd, _ := os.Getwd()
	pos := strings.LastIndex(cwd, "src")
	global.App.ProjectRoot = cwd[:pos]
}

func TestGenIndexYaml(t *testing.T) {
	setup()

	datasource.DefaultGithub.GenIndexYaml()
}

func TestGenArchiveYaml(t *testing.T) {
	setup()

	datasource.DefaultGithub.GenArchiveYaml()
}
