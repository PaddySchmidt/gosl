// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

import "github.com/cpmech/gosl/chk"

func init() {
	Verbose = false
}

func verbose() {
	Verbose = true
	chk.Verbose = true
}
