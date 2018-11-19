// Copyright © 2018 Gavin Chun Jin. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
package combinatorics

import (
	"reflect"
	"testing"
)

var permuteHappyPath = []struct {
	input    string
	expected []string
}{
	{"", []string{}},
	{"A", []string{"A"}},
	{"ST", []string{"ST", "TS"}},
	{"STU", []string{"STU", "SUT", "TSU", "TUS", "UST", "UTS"}},
	{"wxyz", []string{"wxyz", "wxzy", "wyxz", "wyzx", "wzxy", "wzyx", "xwyz", "xwzy", "xywz", "xyzw", "xzwy", "xzyw", "ywxz", "ywzx", "yxwz", "yxzw", "yzwx", "yzxw", "zwxy", "zwyx", "zxwy", "zxyw", "zywx", "zyxw"}},
}

func TestPermute(t *testing.T) {
	for _, tt := range permuteHappyPath {
		reality := Permute(tt.input)
		if !reflect.DeepEqual(reality, tt.expected) {
			t.Errorf("Permute(%#v): expected %#v, actual %#v", tt.input, tt.expected, reality)
		}
	}
}
