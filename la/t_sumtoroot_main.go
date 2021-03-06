// Copyright 2015 Dorival Pedroso and Raul Durand. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/mpi"
)

func setslice(x []float64) {
	switch mpi.Rank() {
	case 0:
		copy(x, []float64{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3})
	case 1:
		copy(x, []float64{10, 10, 10, 20, 20, 20, 30, 30, 30, 40, 40})
	case 2:
		copy(x, []float64{100, 100, 100, 1000, 1000, 1000, 2000, 2000, 2000, 3000, 3000})
	}
}

func main() {

	mpi.Start(false)
	defer func() {
		mpi.Stop(false)
	}()

	if mpi.Rank() == 0 {
		chk.PrintTitle("Test SumToRoot 01")
	}

	M := [][]float64{
		{1000, 1000, 1000, 1011, 1021, 1000},
		{1000, 1000, 1000, 1012, 1022, 1000},
		{1000, 1000, 1000, 1013, 1023, 1000},
		{1011, 1012, 1013, 1000, 1000, 1000},
		{1021, 1022, 1023, 1000, 1000, 1000},
		{1000, 1000, 1000, 1000, 1000, 1000},
	}

	id, sz, m := mpi.Rank(), mpi.Size(), len(M)
	start, endp1 := (id*m)/sz, ((id+1)*m)/sz

	if sz > 6 {
		chk.Panic("this test works with at most 6 processors")
	}

	var J la.Triplet
	J.Init(m, m, m*m)
	for i := start; i < endp1; i++ {
		for j := 0; j < m; j++ {
			J.Put(i, j, M[i][j])
		}
	}
	la.PrintMat(fmt.Sprintf("J @ proc # %d", id), J.ToMatrix(nil).ToDense(), "%10.1f", false)

	la.SpTriSumToRoot(&J)
	var tst testing.T
	if mpi.Rank() == 0 {
		chk.Matrix(&tst, "J @ proc 0", 1.0e-17, J.ToMatrix(nil).ToDense(), [][]float64{
			{1000, 1000, 1000, 1011, 1021, 1000},
			{1000, 1000, 1000, 1012, 1022, 1000},
			{1000, 1000, 1000, 1013, 1023, 1000},
			{1011, 1012, 1013, 1000, 1000, 1000},
			{1021, 1022, 1023, 1000, 1000, 1000},
			{1000, 1000, 1000, 1000, 1000, 1000},
		})
	}
}
