package myutil_test

import (
	"fmt"
	"testing"

	"github.com/marksauter/myutil-go"
	"github.com/stretchr/testify/assert"
)

func TestFromMaybeI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		defaultValue int
		s            *int
		ret          int
	}{
		{0, nil, 0},
		{0, myutil.JustI(1), 1},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("from maybe int %d", test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.FromMaybeI(test.defaultValue, test.s)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestCopyI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s   *int
		ret *int
	}{
		{nil, nil},
		{myutil.JustI(1), myutil.JustI(1)},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("copy maybe int %v", test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.CopyI(test.s)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestHeadMayI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		vs  []int
		ret *int
	}{
		{[]int{}, nil},
		{[]int{1}, myutil.JustI(1)},
		{[]int{2, 3}, myutil.JustI(2)},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%d head", test.vs),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.HeadMayI(test.vs)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestTailMayI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		vs  []int
		ret *int
	}{
		{[]int{}, nil},
		{[]int{1}, myutil.JustI(1)},
		{[]int{2, 3}, myutil.JustI(3)},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%d tail", test.vs),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.TailMayI(test.vs)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestIndexI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		vs  []int
		t   int
		ret int
	}{
		{[]int{1}, 2, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 2}, 2, 1},
		{[]int{1, 2, 3}, 3, 2},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%d? %v", test.t, test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.IndexI(test.vs, test.t)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestHasI(t *testing.T) {
	t.Parallel()

	vs := []int{
		1,
		2,
		3,
	}
	tests := []struct {
		e   int
		has bool
	}{
		{3, true},
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, false},
		{6, false},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%d? %v", test.e, test.has),
			func(t *testing.T) {
				expected := test.has
				actual := myutil.HasI(vs, test.e)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestAnyI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		vs    []int
		f     func(int) bool
		ret   bool
	}{
		{"any equal to",
			[]int{1}, func(v int) bool { return v == 2 }, false},
		{"any equal to",
			[]int{1, 2}, func(v int) bool { return v == 2 }, true},
		{"any > 3",
			[]int{1, 2, 3}, func(v int) bool { return v > 3 }, false},
		{"any > 3",
			[]int{1, 2, 4}, func(v int) bool { return v > 3 }, true},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s? %v", test.label, test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.AnyI(test.vs, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestAllI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		vs    []int
		f     func(int) bool
		ret   bool
	}{
		{"all equal to",
			[]int{1, 2}, func(v int) bool { return v == 1 }, false},
		{"all equal to",
			[]int{1, 1}, func(v int) bool { return v == 1 }, true},
		{"all > 3",
			[]int{1, 2, 3}, func(v int) bool { return v > 3 }, false},
		{"all > 3",
			[]int{4, 5}, func(v int) bool { return v > 3 }, true},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s? %v", test.label, test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.AllI(test.vs, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestFilterI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		vs    []int
		f     func(int) bool
		ret   []int
	}{
		{"filter equal to",
			[]int{1, 2},
			func(v int) bool { return v == 1 },
			[]int{1},
		},
		{"filter equal to",
			[]int{1, 1},
			func(v int) bool { return v == 2 },
			[]int{},
		},
		{"filter > 3",
			[]int{1, 2, 4},
			func(v int) bool { return v > 3 },
			[]int{4},
		},
		{"filter > 3",
			[]int{4, 5},
			func(v int) bool { return v > 3 },
			[]int{4, 5},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s?", test.label),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.FilterI(test.vs, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestMapI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		vs    []int
		f     func(int) int
		ret   []int
	}{
		{"map add suffix",
			[]int{1, 2},
			func(v int) int { return v + 2 },
			[]int{3, 4},
		},
		{"map add prefix",
			[]int{2, 3},
			func(v int) int { return 1 + v },
			[]int{3, 4},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s?", test.label),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.MapI(test.vs, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
				assert.Equal(t, len(test.vs), len(actual), "unexpected return length")
			},
		)
	}
}

func TestApplyI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		v     *int
		f     func(int) int
		ret   *int
	}{
		{"apply to not nil value",
			myutil.JustI(1),
			func(v int) int { return v + 2 },
			myutil.JustI(3),
		},
		{"apply to nil value",
			nil,
			func(v int) int { return v + 2 },
			nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s?", test.label),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.ApplyI(test.v, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}
