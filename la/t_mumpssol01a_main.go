// Copyright 2015 Dorival Pedroso and Raul Durand. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/mpi"
)

func main() {

	mpi.Start(false)
	defer func() {
		mpi.Stop(false)
	}()

	myrank := mpi.Rank()
	if myrank == 0 {
		chk.PrintTitle("Test MUMPS Sol 01a")
	}

	var t la.Triplet
	switch mpi.Size() {
	case 1:
		t.Init(5, 5, 13)
		t.Put(0, 0, 1.0)
		t.Put(0, 0, 1.0)
		t.Put(1, 0, 3.0)
		t.Put(0, 1, 3.0)
		t.Put(2, 1, -1.0)
		t.Put(4, 1, 4.0)
		t.Put(1, 2, 4.0)
		t.Put(2, 2, -3.0)
		t.Put(3, 2, 1.0)
		t.Put(4, 2, 2.0)
		t.Put(2, 3, 2.0)
		t.Put(1, 4, 6.0)
		t.Put(4, 4, 1.0)
	case 2:
		if myrank == 0 {
			t.Init(5, 5, 6)
			t.Put(0, 0, 1.0)
			t.Put(0, 0, 1.0)
			t.Put(1, 0, 3.0)
			t.Put(0, 1, 3.0)
			t.Put(2, 1, -1.0)
			t.Put(4, 1, 4.0)
		} else {
			t.Init(5, 5, 7)
			t.Put(1, 2, 4.0)
			t.Put(2, 2, -3.0)
			t.Put(3, 2, 1.0)
			t.Put(4, 2, 2.0)
			t.Put(2, 3, 2.0)
			t.Put(1, 4, 6.0)
			t.Put(4, 4, 1.0)
		}
	default:
		chk.Panic("this test needs 1 or 2 procs")
	}

	b := []float64{8.0, 45.0, -3.0, 3.0, 19.0}
	x_correct := []float64{1, 2, 3, 4, 5}
	sum_b_to_root := false
	la.RunMumpsTestR(&t, 1e-14, b, x_correct, sum_b_to_root)
}
