package lib

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchMaxSumSubSlice(t *testing.T) {
	testCases := map[string]struct {
		slice  []uint
		result uint
	}{
		"OK 1": {
			slice:  []uint{1},
			result: 1,
		},
		"OK 2": {
			slice:  []uint{1, 2},
			result: 3,
		},
		"OK 3": {
			slice:  []uint{10, 10, 10, 5, 5, 3},
			result: 40,
		},
		"OK 4": {
			slice:  []uint{10, 10, 3, 5, 5},
			result: 23,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(i, func(t *testing.T) {
			t.Parallel()

			result := SearchMaxSumSubSlice(tc.slice)

			require.Equal(t, tc.result, result)
		})
	}
}
