FROM golang:1.23 AS builder

COPY . /go/src/prometheus
WORKDIR /go/src/prometheus/promql

# Write a native Go fuzz wrapper - tagged gofuzz so it can call FuzzParseExpr
RUN cat > fuzz_native_test.go << GOEOF
//go:build gofuzz

package promql

import "testing"

func FuzzParseExprNative(f *testing.F) {
    f.Add([]byte("sum(metric)"))
    f.Fuzz(func(t *testing.T, data []byte) {
        FuzzParseExpr(data)
    })
}
GOEOF

RUN go test -tags gofuzz -c -fuzz . -o /fuzz_prometheus

FROM golang:1.23
COPY --from=builder /fuzz_prometheus /
CMD ["/fuzz_prometheus", "-test.fuzz=FuzzParseExprNative", "-test.fuzzcachedir=/tmp/fuzz"]
