// Copyright © 2018 Gavin Chun Jin. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
package combinatorics

import (
	"fmt"
	"sort"
	"strings"
)

func Permute(input string) []string {
	//fmt.Printf("Permute(%q)\n", input)
	if len(input) >= 1 {
		p := input[len(input)-1]
		result := Permute(input[:len(input)-1])
		//fmt.Printf("\tInsert %q into %#v\n", p, result)
		if len(result) == 0 {
			return []string{string(p)}
		} else {
			tmp := make([]string, 0)
			insertChar := string(p)
			for _, item := range result {
				k := strings.Split(item, "")
				if len(k) == 1 {
					tmp = append(tmp, insertChar+k[0])
					tmp = append(tmp, k[0]+insertChar)
				} else if len(k) > 1 {
					for i, _ := range k {
						//fmt.Printf("\t\t\t\t%d -> %q\n", i, v)
						if i <= len(k) {
							var somestring string
							somestring = strings.Join((k[0:i]), "") + insertChar + strings.Join(k[i:], "")
							fmt.Println(somestring)
							tmp = append(tmp, somestring)
						}
					}
					tmp = append(tmp, item+insertChar)
				}
				//fmt.Printf("\t\tsplit = %v (tmp=%#v)\n", k, tmp)
			}
			sort.Strings(tmp)
			return tmp
		}
	}

	return make([]string, 0)
}
