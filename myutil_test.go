package myutil_test

import (
	"fmt"
	"testing"

	myutil "github.com/marksauter/myutil-go"
	"github.com/stretchr/testify/assert"
)

func TestSlog(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v   interface{}
		ret string
	}{
		{nil, "<nil>"},
		{"foo", "foo"},
		{1, "1"},
		{myutil.JustS("foo"), "foo"},
		{myutil.JustI(1), "1"},
		{[]string{"foo", "bar"}, "[foo bar]"},
		{[]*string{myutil.JustS("foo"), myutil.JustS("bar"), nil}, "[foo bar <nil>]"},
		{[]int{1, 2}, "[1 2]"},
		{[]*int{myutil.JustI(1), myutil.JustI(2), nil}, "[1 2 <nil>]"},
		{map[interface{}]interface{}{"foo": 1}, "map[foo:1]"},
		{map[interface{}]interface{}{myutil.JustS("bar"): myutil.JustI(2)}, "map[bar:2]"},
	}

	for _, test := range tests {
		test := test

		t.Run(
			fmt.Sprintf("sprint maybe string %s", test.ret),
			func(t *testing.T) {
				expected := test.ret
				actual := myutil.Slog(test.v)

				assert.Equal(t, expected, actual, "unexpected return")
			},
		)
	}
}
