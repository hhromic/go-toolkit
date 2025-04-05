// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"strconv"
	"testing"

	tktypes "github.com/hhromic/go-toolkit/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			want: 2,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			assert.Equal(t, tCase.want, tCase.ranges.Len())
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
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			i:         0,
			j:         1,
			want:      true,
			wantPanic: false,
		},
		{
			name: "ThreeElementsReversed",
			ranges: tktypes.Ranges{
				{Min: 5, Max: 7, Value: "bar"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 1, Max: 3, Value: "foo"},
			},
			i:         0,
			j:         1,
			want:      false,
			wantPanic: false,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			if tCase.wantPanic {
				assert.Panics(t, func() {
					tCase.ranges.Less(tCase.i, tCase.j)
				})
			} else {
				assert.Equal(t, tCase.want, tCase.ranges.Less(tCase.i, tCase.j))
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
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			i: 0,
			j: 1,
			want: tktypes.Ranges{
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			wantPanic: false,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			if tCase.wantPanic {
				assert.Panics(t, func() {
					tCase.ranges.Swap(tCase.i, tCase.j)
				})
			} else {
				tCase.ranges.Swap(tCase.i, tCase.j)
				assert.Equal(t, tCase.want, tCase.ranges)
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
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
				{Min: 1, Max: 3, Value: "foo"},
			},
			want: tktypes.Ranges{
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			tCase.ranges.Sort()
			assert.Equal(t, tCase.want, tCase.ranges)
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
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			v:    0,
			want: nil,
		},
		{
			name: "ThreeElementsOne",
			ranges: tktypes.Ranges{
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			v:    1,
			want: "foo",
		},
		{
			name: "ThreeElementsTwo",
			ranges: tktypes.Ranges{
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			v:    2,
			want: "foo",
		},
		{
			name: "ThreeElementsFour",
			ranges: tktypes.Ranges{
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			v:    4,
			want: "baz",
		},
		{
			name: "ThreeElementsFive",
			ranges: tktypes.Ranges{
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			v:    5,
			want: "baz",
		},
		{
			name: "ThreeElementsSix",
			ranges: tktypes.Ranges{
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			v:    6,
			want: "bar",
		},
		{
			name: "ThreeElementsEight",
			ranges: tktypes.Ranges{
				{Min: 1, Max: 3, Value: "foo"},
				{Min: 2, Max: 5, Value: "baz"},
				{Min: 5, Max: 7, Value: "bar"},
			},
			v:    8,
			want: nil,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			assert.Equal(t, tCase.want, tCase.ranges.Search(tCase.v))
		})
	}
}

func TestBareRangeMarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		r       tktypes.BareRange
		want    []byte
		wantErr error
	}{
		{
			name:    "OpenLeft",
			r:       tktypes.BareRange{Min: tktypes.RangeMin, Max: 10, Value: struct{}{}},
			want:    []byte(":10"),
			wantErr: nil,
		},
		{
			name:    "Closed",
			r:       tktypes.BareRange{Min: 20, Max: 30, Value: struct{}{}},
			want:    []byte("20:30"),
			wantErr: nil,
		},
		{
			name:    "OpenRight",
			r:       tktypes.BareRange{Min: 40, Max: tktypes.RangeMax, Value: struct{}{}},
			want:    []byte("40:"),
			wantErr: nil,
		},
		{
			name:    "Single",
			r:       tktypes.BareRange{Min: 50, Max: 50, Value: struct{}{}},
			want:    []byte("50"),
			wantErr: nil,
		},
		{
			name: "FullRange",
			r: tktypes.BareRange{
				Min:   tktypes.RangeMin,
				Max:   tktypes.RangeMax,
				Value: struct{}{},
			},
			want:    []byte(":"),
			wantErr: nil,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			b, err := tCase.r.MarshalText()
			require.ErrorIs(t, err, tCase.wantErr)

			if tCase.wantErr == nil {
				assert.Equal(t, tCase.want, b)
			}
		})
	}
}

//nolint:funlen
func TestBareRangeUnmarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		b       []byte
		want    tktypes.BareRange
		wantErr error
	}{
		{
			name:    "OpenLeft",
			b:       []byte(":10"),
			want:    tktypes.BareRange{Min: tktypes.RangeMin, Max: 10, Value: struct{}{}},
			wantErr: nil,
		},
		{
			name:    "Closed",
			b:       []byte("20:30"),
			want:    tktypes.BareRange{Min: 20, Max: 30, Value: struct{}{}},
			wantErr: nil,
		},
		{
			name:    "OpenRight",
			b:       []byte("40:"),
			want:    tktypes.BareRange{Min: 40, Max: tktypes.RangeMax, Value: struct{}{}},
			wantErr: nil,
		},
		{
			name:    "Single",
			b:       []byte("50"),
			want:    tktypes.BareRange{Min: 50, Max: 50, Value: struct{}{}},
			wantErr: nil,
		},
		{
			name: "FullRange",
			b:    []byte(":"),
			want: tktypes.BareRange{
				Min:   tktypes.RangeMin,
				Max:   tktypes.RangeMax,
				Value: struct{}{},
			},
			wantErr: nil,
		},
		{
			name:    "InvalidFormat",
			b:       []byte("foo::bar"),
			want:    tktypes.BareRange{}, //nolint:exhaustruct
			wantErr: tktypes.ErrUnknownFormat,
		},
		{
			name:    "InvalidSyntaxSingle",
			b:       []byte("foo"),
			want:    tktypes.BareRange{}, //nolint:exhaustruct
			wantErr: strconv.ErrSyntax,
		},
		{
			name:    "InvalidSyntaxRange",
			b:       []byte("foo:bar"),
			want:    tktypes.BareRange{}, //nolint:exhaustruct
			wantErr: strconv.ErrSyntax,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			var r tktypes.BareRange
			err := r.UnmarshalText(tCase.b)
			require.ErrorIs(t, err, tCase.wantErr)

			if tCase.wantErr == nil {
				assert.Equal(t, tCase.want, r)
			}
		})
	}
}

func TestBareRangesMarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		ranges  tktypes.BareRanges
		want    []byte
		wantErr error
	}{
		{
			name:    "Empty",
			ranges:  tktypes.BareRanges{},
			want:    []byte(""),
			wantErr: nil,
		},
		{
			name: "OneElement",
			ranges: tktypes.BareRanges{
				{Min: 10, Max: 20, Value: struct{}{}},
			},
			want:    []byte("10:20"),
			wantErr: nil,
		},
		{
			name: "ThreeElements",
			ranges: tktypes.BareRanges{
				{Min: tktypes.RangeMin, Max: 10, Value: struct{}{}},
				{Min: 20, Max: 30, Value: struct{}{}},
				{Min: 40, Max: tktypes.RangeMax, Value: struct{}{}},
			},
			want:    []byte(":10,20:30,40:"),
			wantErr: nil,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			b, err := tCase.ranges.MarshalText()
			require.ErrorIs(t, err, tCase.wantErr)

			if tCase.wantErr == nil {
				assert.Equal(t, tCase.want, b)
			}
		})
	}
}

func TestBareRangesUnmarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		b       []byte
		want    tktypes.BareRanges
		wantErr error
	}{
		{
			name:    "Empty",
			b:       []byte(""),
			want:    tktypes.BareRanges{},
			wantErr: nil,
		},
		{
			name: "OneElement",
			b:    []byte("10:20"),
			want: tktypes.BareRanges{
				{Min: 10, Max: 20, Value: struct{}{}},
			},
			wantErr: nil,
		},
		{
			name: "ThreeElements",
			b:    []byte(":10,20:30,40:"),
			want: tktypes.BareRanges{
				{Min: tktypes.RangeMin, Max: 10, Value: struct{}{}},
				{Min: 20, Max: 30, Value: struct{}{}},
				{Min: 40, Max: tktypes.RangeMax, Value: struct{}{}},
			},
			wantErr: nil,
		},
		{
			name:    "InvalidFormat",
			b:       []byte(":10;20:30;40:"),
			want:    tktypes.BareRanges{},
			wantErr: tktypes.ErrUnknownFormat,
		},
		{
			name:    "InvalidSyntax",
			b:       []byte(":bar,foo:30,40:"),
			want:    tktypes.BareRanges{},
			wantErr: strconv.ErrSyntax,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			var ranges tktypes.BareRanges
			err := ranges.UnmarshalText(tCase.b)
			require.ErrorIs(t, err, tCase.wantErr)

			if tCase.wantErr == nil {
				assert.Equal(t, tCase.want, ranges)
			}
		})
	}
}
