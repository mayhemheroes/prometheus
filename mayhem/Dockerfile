FROM golang:1.23 AS builder

COPY . /go/src/prometheus
WORKDIR /go/src/prometheus/promql

RUN go test -tags gofuzz -c -fuzz . -o /fuzz_prometheus

FROM golang:1.23
COPY --from=builder /fuzz_prometheus /
CMD ["/fuzz_prometheus", "-test.fuzz=FuzzParseExprNative", "-test.fuzzcachedir=/tmp/fuzz"]
