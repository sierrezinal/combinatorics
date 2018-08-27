// Copyright © 2018 Gavin Chun Jin. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
package combinatorics

import (
	"testing"
)

var fixturesHappyPath = []struct {
	n        int
	k        int
	expected int
}{
	{0, 0, 1},
	{1, 0, 1},
	{1, 1, 1},
	{6, 0, 1},
	{6, 1, 6},
	{6, 2, 15},
	{6, 3, 20},
	{6, 4, 15},
	{6, 5, 6},
	{6, 6, 1},
	{10, 2, 45},
	{1000, 2, 499500},
}

func TestHappyPath(t *testing.T) {
	for _, tt := range fixturesHappyPath {
		reality, _ := Nchoosek(tt.n, tt.k)
		if reality != tt.expected {
			t.Errorf("Nchoosek(%v, %v): expected %v, actual %v", tt.n, tt.k, tt.expected, reality)
		}
	}
}

var fixturesSadPath = []struct {
	n        int
	k        int
	expected error
}{
	{-1, 1, ErrNegativeArgument},
	{1, -1, ErrNegativeArgument},
	{-1, -1, ErrNegativeArgument},
	{0, 1, ErrKbiggerThanN},
	{1, 2, ErrKbiggerThanN},
	{1, 1, nil},
}

func TestSadPath(t *testing.T) {
	for _, tt := range fixturesSadPath {
		_, err := Nchoosek(tt.n, tt.k)
		if err != tt.expected {
			t.Errorf("Nchoosek(%v, %v): expected %v, actual %v", tt.n, tt.k, tt.expected, err)
		}
	}
}
