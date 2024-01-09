// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"fmt"

	tktypes "github.com/hhromic/go-toolkit/types"
)

func ExampleRanges_Sort() {
	var ranges tktypes.Ranges

	ranges = append(ranges, tktypes.Range{Min: 4, Max: 4, Value: "cat"})
	ranges = append(ranges, tktypes.Range{Min: 1, Max: 2, Value: "dog"})
	ranges = append(ranges, tktypes.Range{Min: 3, Max: 8, Value: "fox"})

	ranges.Sort()
	fmt.Println(ranges)

	// Output:
	// [{1 2 dog} {3 8 fox} {4 4 cat}]
}

func ExampleRanges_Search() {
	var ranges tktypes.Ranges

	ranges = append(ranges, tktypes.Range{Min: 1, Max: 2, Value: "dog"})
	ranges = append(ranges, tktypes.Range{Min: 4, Max: 4, Value: "cat"})

	for i := 0; i < 6; i++ {
		fmt.Printf("value %d belongs to %v\n", i, ranges.Search(i))
	}

	// Output:
	// value 0 belongs to <nil>
	// value 1 belongs to dog
	// value 2 belongs to dog
	// value 3 belongs to <nil>
	// value 4 belongs to cat
	// value 5 belongs to <nil>
}
