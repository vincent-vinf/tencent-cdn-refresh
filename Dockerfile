FROM golang:1.18 as builder
WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 go build -o bin/tencent-cdn-refresh ./main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/bin/* /usr/local/bin
USER root
