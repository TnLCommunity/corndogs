FROM golang:1.20 as builder

WORKDIR /workspace
# install grpc health probe
ENV GRPC_HEALTH_PROBE_VERSION=v0.4.18
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && chmod +x /bin/grpc_health_probe
# copy relevant things
COPY go.mod .
COPY go.sum .
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download
COPY server/ ./server
COPY cmd/ ./cmd
COPY main.go ./main.go
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o main main.go

# Use distroless as minimal base image
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/main .
COPY --from=builder /bin/grpc_health_probe ./grpc_health_probe
USER 65532:65532

ENTRYPOINT ["/main"]
CMD ["run"]
