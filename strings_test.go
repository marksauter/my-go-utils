package myutil_test

import (
	"fmt"
	"testing"

	"github.com/marksauter/myutil-go"
	"github.com/stretchr/testify/assert"
)

func TestHeadMayS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ss  []string
		ret *string
	}{
		{[]string{}, nil},
		{[]string{"foo"}, myutil.NewS("foo")},
		{[]string{"bar", "baz"}, myutil.NewS("bar")},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s head", test.ss),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.HeadMayS(test.ss)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestTailMayS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ss  []string
		ret *string
	}{
		{[]string{}, nil},
		{[]string{"foo"}, myutil.NewS("foo")},
		{[]string{"bar", "baz"}, myutil.NewS("baz")},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s tail", test.ss),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.TailMayS(test.ss)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestJustS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s   *string
		ret string
	}{
		{nil, ""},
		{myutil.NewS("foo"), "foo"},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("force string %s", test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.JustS(test.s)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestIndexS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		vs  []string
		t   string
		ret int
	}{
		{[]string{"foo"}, "bar", -1},
		{[]string{"foo"}, "foo", 0},
		{[]string{"foo", "bar"}, "bar", 1},
		{[]string{"foo", "bar", "baz"}, "baz", 2},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s? %v", test.t, test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.IndexS(test.vs, test.t)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestHasS(t *testing.T) {
	t.Parallel()

	ss := []string{
		"foo",
		"bar",
		"baz",
	}
	tests := []struct {
		e   string
		has bool
	}{
		{"foobar", false},
		{"foo", true},
		{"bar", true},
		{"baz", true},
		{"qux", false},
		{"quux", false},
		{"quuz", false},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s? %v", test.e, test.has),
			func(t *testing.T) {
				expected := test.has
				actual := myutil.HasS(ss, test.e)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestAnyS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		vs    []string
		f     func(string) bool
		ret   bool
	}{
		{"any equal to",
			[]string{"foo"}, func(s string) bool { return s == "bar" }, false},
		{"any equal to",
			[]string{"foo", "bar"}, func(s string) bool { return s == "bar" }, true},
		{"any len > 3",
			[]string{"foo", "bar", "baz"}, func(s string) bool { return len(s) > 3 }, false},
		{"any len > 3",
			[]string{"foo", "bar", "quux"}, func(s string) bool { return len(s) > 3 }, true},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s? %v", test.label, test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.AnyS(test.vs, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestAllS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		vs    []string
		f     func(string) bool
		ret   bool
	}{
		{"all equal to",
			[]string{"foo", "bar"}, func(s string) bool { return s == "foo" }, false},
		{"all equal to",
			[]string{"foo", "foo"}, func(s string) bool { return s == "foo" }, true},
		{"all len > 3",
			[]string{"foo", "bar", "quux"}, func(s string) bool { return len(s) > 3 }, false},
		{"all len > 3",
			[]string{"quux", "quuz"}, func(s string) bool { return len(s) > 3 }, true},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s? %v", test.label, test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.AllS(test.vs, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestFilterS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		vs    []string
		f     func(string) bool
		ret   []string
	}{
		{"filter equal to",
			[]string{"foo", "bar"},
			func(s string) bool { return s == "foo" },
			[]string{"foo"},
		},
		{"filter equal to",
			[]string{"foo", "foo"},
			func(s string) bool { return s == "bar" },
			[]string{},
		},
		{"filter len > 3",
			[]string{"foo", "bar", "quux"},
			func(s string) bool { return len(s) > 3 },
			[]string{"quux"},
		},
		{"filter len > 3",
			[]string{"quux", "quuz"},
			func(s string) bool { return len(s) > 3 },
			[]string{"quux", "quuz"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s?", test.label),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.FilterS(test.vs, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestMapS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		vs    []string
		f     func(string) string
		ret   []string
	}{
		{"map add suffix",
			[]string{"foo", "bar"},
			func(s string) string { return s + "bar" },
			[]string{"foobar", "barbar"},
		},
		{"map add prefix",
			[]string{"bar", "baz"},
			func(s string) string { return "foo" + s },
			[]string{"foobar", "foobaz"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s?", test.label),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.MapS(test.vs, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
				assert.Equal(t, len(test.vs), len(actual), "unexpected return length")
			},
		)
	}
}

func TestApplyS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		label string
		v     *string
		f     func(string) string
		ret   *string
	}{
		{"apply to not nil value",
			myutil.NewS("foo"),
			func(s string) string { return s + "bar" },
			myutil.NewS("foobar"),
		},
		{"apply to nil value",
			nil,
			func(s string) string { return s + "bar" },
			nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s?", test.label),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.ApplyS(test.v, test.f)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}
