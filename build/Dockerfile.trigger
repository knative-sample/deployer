# Build the manager binary
FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/golang:1.12 as builder

# Copy in the go src
WORKDIR /go/src/github.com/knative-sample/deployer/
COPY cmd/   cmd/
COPY pkg/   pkg/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/trigger github.com/knative-sample/deployer/cmd/trigger/

FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/alpine:3.9
WORKDIR /app/
RUN mkdir -p /app/bin/
COPY --from=builder /go/src/github.com/knative-sample/deployer/bin/trigger /app/bin/trigger
ENTRYPOINT ["/app/bin/trigger"]
