# server
FROM golang:1.13.1-alpine as builder

ENV GO111MODULE=on
#ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o server cmd/server/main.go
RUN go get -u github.com/CatchZeng/jenkinsapi

# web
FROM node:8.12-alpine as web-builder
#RUN npm install -g nrm@1.2.1 && nrm use taobao
ADD ./web/appboot /app
WORKDIR /app
RUN npm install && npm run build

FROM alpine:3.7

# https://pkgs.alpinelinux.org/packages
RUN apk add  --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/v3.7/main/ git=2.15.4-r0
RUN apk add  --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/v3.7/main/ bash=4.4.19-r1
RUN apk add  --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/v3.7/main/ openssh=7.5_p1-r10
RUN apk add  --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/v3.7/main/ nginx=1.12.2-r4

RUN mkdir /server
RUN mkdir /app

# jenkinsapi
COPY --from=builder /go/bin/jenkinsapi /usr/bin/jenkinsapi

# appboot server
COPY --from=builder /build/server /server

# appboot web
COPY --from=web-builder /app/dist /app
COPY --from=web-builder /app/start.sh /app
COPY --from=web-builder /app/nginx.conf /etc/nginx/nginx.conf

# start script
ADD build/package/start.sh /

WORKDIR /

EXPOSE 8000

CMD ["./start.sh"]