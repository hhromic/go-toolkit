// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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

// BareRange is an alias for marshaling/unmarshaling bare ranges with no significant values.
type BareRange = Range

// MarshalText implements [encoding.TextMarshaler] for a bare range.
// The output format is "min:max". If min or max are [math.MinInt] or [math.MaxInt] respectively,
// their values are omitted in the output: "min:" or ":max".
func (r BareRange) MarshalText() ([]byte, error) {
	var out string

	switch {
	case r.Min == math.MinInt && r.Max != math.MaxInt:
		out = fmt.Sprintf(":%d", r.Max)
	case r.Min != math.MinInt && r.Max == math.MaxInt:
		out = fmt.Sprintf("%d:", r.Min)
	default:
		out = fmt.Sprintf("%d:%d", r.Min, r.Max)
	}

	return []byte(out), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler].
// It accepts any slice of bytes produced by [BareRange.MarshalText].
func (r *BareRange) UnmarshalText(b []byte) error {
	str := string(b)
	switch strings.Count(str, ":") {
	case 0:
		n, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			return fmt.Errorf("%q: parse int: %w", str, err)
		}

		*r = Range{Min: int(n), Max: int(n), Value: struct{}{}}
	case 1:
		var min, max int64

		parts := strings.SplitN(str, ":", 2) //nolint:gomnd
		if parts[0] == "" {
			min = math.MinInt
		} else {
			var err error
			if min, err = strconv.ParseInt(parts[0], 10, 0); err != nil {
				return fmt.Errorf("%q: parse int: %w", parts[0], err)
			}
		}

		if parts[1] == "" {
			max = math.MaxInt
		} else {
			var err error
			if max, err = strconv.ParseInt(parts[1], 10, 0); err != nil {
				return fmt.Errorf("%q: parse int: %w", parts[1], err)
			}
		}

		*r = Range{Min: int(min), Max: int(max), Value: struct{}{}}
	default:
		return fmt.Errorf("%q: %w", str, ErrUnknownFormat)
	}

	return nil
}
