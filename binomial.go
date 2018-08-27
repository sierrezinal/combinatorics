// Copyright © 2018 Gavin Chun Jin. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
package combinatorics

import (
	"errors"
)

// ErrKbiggerThanN means k is outside of [0,n]
var ErrKbiggerThanN = errors.New("k does not belong to range 0 <= k <=n")

// ErrNegativeArgument means that either one of n or k is negative.
var ErrNegativeArgument = errors.New("Negative arguments")

// Nchoosek computes the Binomial coefficients.
// Both arguments k, n must be positive integers.
// 0 <= k <= n
func Nchoosek(n, k int) (int, error) {
	if n < 0 || k < 0 {
		return 0, ErrNegativeArgument
	}
	if k > n {
		return 0, ErrKbiggerThanN
	}
	if k == 0 {
		return 1, nil
	}

	var array [][]int

	array = make([][]int, n+1)
	for i, _ := range array {
		array[i] = make([]int, k+1)
		array[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			array[i][j] = array[i-1][j-1] + array[i-1][j] // Pascal's triangle
		}
	}

	return array[n][k], nil
}
