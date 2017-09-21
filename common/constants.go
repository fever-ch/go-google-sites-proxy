// Copyright 2017 Fever.ch Authors. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package common

type ProgramInfoStruct struct {
	Name       string
	Fullname   string
	Version    string
	Git        string
	BuildDate  string
	ProjectUrl string
}

var ProgramInfo = ProgramInfoStruct{"GSSP",
	"Go Google Site Proxy",
	"0.0.1-SNAPSHOT",
	"undefined",
	"undefined",
	"https://ggsp.fever.ch"}
