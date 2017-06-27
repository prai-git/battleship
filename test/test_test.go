// Copyright 2015 The draw2d Authors. All rights reserved.
// created: 26/06/2015 by Stani Michiels

// Package draw2dpdf_test gives test coverage with the command:
// go test -cover ./... | grep -v "no test"
// (It should be run from its parent draw2d directory.)
package battleship_test

import (
	"testing"
)

type sample func() error

func test(t *testing.T, battleship sample) {

	err := battleship()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
