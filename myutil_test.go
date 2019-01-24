package myutil_test

import (
	"fmt"
	"testing"

	"github.com/marksauter/myutil"
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
