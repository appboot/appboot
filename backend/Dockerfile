FROM golang:1.13.1-alpine as builder
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o main main.go

FROM alpine:3.7
RUN mkdir /app
COPY --from=builder /build/main /app
WORKDIR /app
EXPOSE 8888
CMD ["./main"]