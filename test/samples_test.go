// Copyright 2015 The draw2d Authors. All rights reserved.
// created: 26/06/2015 by Stani Michiels
// See also test_test.go

package battleship_test

import (
	"testing"

	"github.com/prai-git/battleship/test/samples/test1"
	"github.com/prai-git/battleship/test/samples/test2"
)

func TestSampleTest1(t *testing.T) {
	test(t, test1.Main)
}

func TestSampleTest2(t *testing.T) {
	test(t, test2.Main)
}
