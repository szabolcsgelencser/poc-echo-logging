package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var pocCases = map[string]struct {
	a, b    int
	want    int
	wantErr error
}{
	"my first case":  {a: 1, b: 2, want: 3},
	"my second case": {a: 5, b: 6, want: 11},
	"errorous case":  {a: -1, b: -2, wantErr: errNegative},
}

func TestPOC(t *testing.T) {
	for name, tc := range pocCases {
		t.Run(name, func(t *testing.T) {
			got, gotErr := sum(tc.a, tc.b)
			if tc.wantErr != nil {
				require.Equal(t, tc.wantErr, gotErr)
				return
			}
			require.NoError(t, gotErr)
			require.Equal(t, tc.want, got)
		})
	}
}
