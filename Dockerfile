FROM golang:1.22.3 as builder
WORKDIR /code
COPY . .
RUN go env -w GO111MODULE=on &&  \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go build -o   ./server  ./cmd/server/

FROM alpine:3.17 as final
WORKDIR /run
COPY --from=builder /code/server /run/server
RUN chmod +x /run/server
EXPOSE 8080
CMD ["/run/server"]