// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"sort"
)

// Range is a min/max range (inclusive) of integers that references a value of any type.
//
// Source: https://stackoverflow.com/a/39750394
type Range struct {
	Min, Max int
	Value    any
}

// Ranges is a collection of sortable and searchable [Range] instances.
// It implements the [sort.Interface] interface.
//
// Source: https://stackoverflow.com/a/39750394
type Ranges []Range

// Len is the number of [Range] elements in the collection.
//
// Source: https://stackoverflow.com/a/39750394
func (r Ranges) Len() int {
	return len(r)
}

// Less reports whether the [Range] element with index i
// must sort before the [Range] element with index j.
//
// Source: https://stackoverflow.com/a/39750394
func (r Ranges) Less(i, j int) bool {
	return r[i].Min < r[j].Min
}

// Swap swaps the [Range] elements with indexes i and j.
//
// Source: https://stackoverflow.com/a/39750394
func (r Ranges) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Sort sorts the collection in ascending order as determined by the [Ranges.Less] method.
//
// Source: https://stackoverflow.com/a/39750394
func (r Ranges) Sort() {
	sort.Sort(r)
}

// Search uses binary search to find and return the first [Range] element in the collection
// in which v is contained (min/max range values are inclusive).
// This function uses the [sort.Search] function.
//
// Source: https://stackoverflow.com/a/39750394
func (r Ranges) Search(v int) any {
	ln := r.Len()
	if i := sort.Search(ln, func(i int) bool { return v <= r[i].Max }); i < ln {
		if it := &r[i]; v >= it.Min && v <= it.Max {
			return it.Value
		}
	}

	return nil
}
