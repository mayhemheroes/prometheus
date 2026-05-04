//go:build gofuzz

package promql

import "testing"

func FuzzParseExprNative(f *testing.F) {
	f.Add([]byte("sum(metric)"))
	f.Fuzz(func(t *testing.T, data []byte) {
		FuzzParseExpr(data)
	})
}
