// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"testing"

	tktypes "github.com/hhromic/go-toolkit/types"
	"github.com/stretchr/testify/assert"
)

func TestRangesLen(t *testing.T) {
	testCases := []struct {
		name   string
		ranges tktypes.Ranges
		want   int
	}{
		{
			name:   "Empty",
			ranges: tktypes.Ranges{},
			want:   0,
		},
		{
			name: "TwoElements",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.ranges.Len())
		})
	}
}

func TestRangesLess(t *testing.T) {
	testCases := []struct {
		name      string
		ranges    tktypes.Ranges
		i, j      int
		want      bool
		wantPanic bool
	}{
		{
			name:      "Empty",
			ranges:    tktypes.Ranges{},
			i:         0,
			j:         1,
			want:      false,
			wantPanic: true,
		},
		{
			name: "ThreeElements",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			i:         0,
			j:         1,
			want:      true,
			wantPanic: false,
		},
		{
			name: "ThreeElementsReversed",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
			},
			i:         0,
			j:         1,
			want:      false,
			wantPanic: false,
		},
	}

	for _, tc := range testCases { //nolint:varnamelen
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantPanic {
				assert.Panics(t, func() {
					tc.ranges.Less(tc.i, tc.j)
				})
			} else {
				assert.Equal(t, tc.want, tc.ranges.Less(tc.i, tc.j))
			}
		})
	}
}

func TestRangesSwap(t *testing.T) {
	testCases := []struct {
		name      string
		ranges    tktypes.Ranges
		i, j      int
		want      tktypes.Ranges
		wantPanic bool
	}{
		{
			name:      "Empty",
			ranges:    tktypes.Ranges{},
			i:         0,
			j:         1,
			want:      tktypes.Ranges{},
			wantPanic: true,
		},
		{
			name: "ThreeElements",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			i: 0,
			j: 1,
			want: tktypes.Ranges{
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			wantPanic: false,
		},
	}

	for _, tc := range testCases { //nolint:varnamelen
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantPanic {
				assert.Panics(t, func() {
					tc.ranges.Swap(tc.i, tc.j)
				})
			} else {
				tc.ranges.Swap(tc.i, tc.j)
				assert.Equal(t, tc.want, tc.ranges)
			}
		})
	}
}

func TestRangesSort(t *testing.T) {
	testCases := []struct {
		name   string
		ranges tktypes.Ranges
		want   tktypes.Ranges
	}{
		{
			name:   "Empty",
			ranges: tktypes.Ranges{},
			want:   tktypes.Ranges{},
		},
		{
			name: "ThreeElements",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
			},
			want: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.ranges.Sort()
			assert.Equal(t, tc.want, tc.ranges)
		})
	}
}

//nolint:funlen
func TestRangesSearch(t *testing.T) {
	testCases := []struct {
		name   string
		ranges tktypes.Ranges
		v      int
		want   any
	}{
		{
			name:   "Empty",
			ranges: tktypes.Ranges{},
			v:      10,
			want:   nil,
		},
		{
			name: "ThreeElementsNotFound",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			v:    0,
			want: nil,
		},
		{
			name: "ThreeElementsOne",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			v:    1,
			want: "foo",
		},
		{
			name: "ThreeElementsTwo",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			v:    2,
			want: "foo",
		},
		{
			name: "ThreeElementsFour",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			v:    4,
			want: "baz",
		},
		{
			name: "ThreeElementsFive",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			v:    5,
			want: "baz",
		},
		{
			name: "ThreeElementsSix",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			v:    6,
			want: "bar",
		},
		{
			name: "ThreeElementsEight",
			ranges: tktypes.Ranges{
				tktypes.Range{Min: 1, Max: 3, Value: "foo"},
				tktypes.Range{Min: 2, Max: 5, Value: "baz"},
				tktypes.Range{Min: 5, Max: 7, Value: "bar"},
			},
			v:    8,
			want: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.ranges.Search(tc.v))
		})
	}
}