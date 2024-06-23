FROM docker.signit.cn/base/golang:1.22.3 as builder
WORKDIR /demo-code
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN   go mod tidy
RUN  CGO_ENABLED=0 go build -o   /demo-code/server  ./cmd/server/

FROM docker.signit.cn/base/alpine:3.17
COPY --from=builder /demo-code/config.yaml /config.yaml
COPY --from=builder /demo-code/server /server
EXPOSE 8080
CMD ["/server"]