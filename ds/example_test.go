// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package ds_test

import (
	"fmt"

	"github.com/hhromic/go-toolkit/ds"
)

func ExampleRanges_Sort() {
	ranges := ds.Ranges{
		{Min: 4, Max: 4, Value: "cat"},
		{Min: 1, Max: 2, Value: "dog"},
		{Min: 3, Max: 8, Value: "fox"},
	}

	ranges.Sort()
	fmt.Println(ranges)

	// Output:
	// [{1 2 dog} {3 8 fox} {4 4 cat}]
}

func ExampleRanges_Search() {
	ranges := ds.Ranges{
		{Min: 1, Max: 2, Value: "dog"},
		{Min: 4, Max: 4, Value: "cat"},
	}

	for i := range 6 {
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

func ExampleBareRange_MarshalText() {
	ranges := []ds.BareRange{
		{Min: ds.RangeMin, Max: 10, Value: struct{}{}},
		{Min: 20, Max: 30, Value: struct{}{}},
		{Min: 40, Max: ds.RangeMax, Value: struct{}{}},
		{Min: 50, Max: 50, Value: struct{}{}},
		{Min: ds.RangeMin, Max: ds.RangeMax, Value: struct{}{}},
	}

	for _, r := range ranges {
		b, err := r.MarshalText()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))
	}

	// Output:
	// :10
	// 20:30
	// 40:
	// 50
	// :
}

func ExampleBareRange_UnmarshalText() {
	for _, t := range []string{":10", "20:30", "40:", "50", ":"} {
		var rng ds.BareRange

		err := rng.UnmarshalText([]byte(t))
		if err != nil {
			panic(err)
		}

		fmt.Println(rng)
	}

	// Output:
	// {-9223372036854775808 10 {}}
	// {20 30 {}}
	// {40 9223372036854775807 {}}
	// {50 50 {}}
	// {-9223372036854775808 9223372036854775807 {}}
}

func ExampleBareRanges_MarshalText() {
	ranges := ds.BareRanges{
		{Min: ds.RangeMin, Max: 10, Value: struct{}{}},
		{Min: 20, Max: 30, Value: struct{}{}},
		{Min: 40, Max: ds.RangeMax, Value: struct{}{}},
		{Min: 50, Max: 50, Value: struct{}{}},
		{Min: ds.RangeMin, Max: ds.RangeMax, Value: struct{}{}},
	}

	b, err := ranges.MarshalText()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	// Output:
	// :10,20:30,40:,50,:
}

//nolint:lll
func ExampleBareRanges_UnmarshalText() {
	t := ":10,20:30,40:,50,:"

	var rngs ds.BareRanges

	err := rngs.UnmarshalText([]byte(t))
	if err != nil {
		panic(err)
	}

	fmt.Println(rngs)

	// Output:
	// [{-9223372036854775808 10 {}} {20 30 {}} {40 9223372036854775807 {}} {50 50 {}} {-9223372036854775808 9223372036854775807 {}}]
}
