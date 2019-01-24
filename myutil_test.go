package myutil_test

import (
	"fmt"
	"testing"

	"github.com/marksauter/my-go-utils"
	"github.com/stretchr/testify/assert"
)

func TestElemString(t *testing.T) {
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
				actual := myutil.ElemString(test.e, ss)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}

func TestHeadStringMay(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ss  []string
		ret *string
	}{
		{[]string{}, nil},
		{[]string{"foo"}, myutil.NewString("foo")},
		{[]string{"bar", "baz"}, myutil.NewString("bar")},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("%s head", test.ss),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.HeadStringMay(test.ss)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}
