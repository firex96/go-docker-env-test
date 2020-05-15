# Build binary
FROM golang:1.14 AS builder
WORKDIR /build

ENV CGO_ENABLED=0

COPY go.mod ./
RUN go mod download

COPY . ./
RUN go build -gcflags "all=-N -l" . # -gcflags "all=-N -l"로 컴파일러 최적화 옵션을 끈다.

# Install delve
RUN go get -u github.com/go-delve/delve/cmd/dlv # delve 설치

# Run on scratch
FROM alpine:3.10
WORKDIR /bin

COPY --from=builder /build/calc /bin/
COPY --from=builder /go/bin/dlv /bin/ # delve 가져오기

EXPOSE 8000
EXPOSE 2345 # 디버거 서버 포트
ENTRYPOINT ["/bin/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/bin/calc", "--continue"] # 디버깅 시작