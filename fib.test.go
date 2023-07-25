package main

import (
	fib "FibaTest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func main() {
	TestFib(*testing.T{})
}
func TestFib(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "one",
			n:    1,
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fib.Fib(tc.n))

		})
	}
}
