// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rnd

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

func Test_hc01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("hc01. hypercube")

	Init(111)

	n := 10
	x := LatinIHS(2, n, 5)
	io.Pforan("x = %v\n", x)

	xcor := [][]int{
		{2, 9, 5, 8, 1, 4, 3, 10, 7, 6},
		{3, 10, 1, 2, 7, 4, 9, 6, 5, 8},
	}
	chk.IntMat(tst, "x", x, xcor)

	if chk.Verbose {
		xrange := [][]float64{
			{-1.0, 1.0},
			{0.0, 2.0},
		}
		PlotHc2d("/tmp/gosl", "test_hc01", x, xrange)
	}
}

func Test_hc02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("hc02. hypercube")

	Init(0)

	n := 36
	x := LatinIHS(2, n, 5)
	io.Pforan("x = %v\n", x)

	if chk.Verbose {
		xrange := [][]float64{
			{-1.0, 1.0},
			{0.0, 2.0},
		}
		PlotHc2d("/tmp/gosl", "test_hc02", x, xrange)
	}
}

func Test_hc03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("hc03. hypercube (3D)")

	Init(0)

	n := 100
	x := LatinIHS(3, n, 5)
	io.Pforan("x = %v\n", x)

	if chk.Verbose {
		xrange := [][]float64{
			{-1.0, 1.0},
			{0.0, 2.0},
			{-2.0, 0.0},
		}
		PlotHc3d("/tmp/gosl", "test_hc03", x, xrange, false)
	}
}
