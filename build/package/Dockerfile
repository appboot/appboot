# server
FROM golang:1.16.4-alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy && go build -o server cmd/server/main.go
RUN go get -u github.com/CatchZeng/jenkinsapi

# web
FROM node:8.12-alpine as web-builder
#RUN npm install -g nrm@1.2.1 && nrm use taobao
ADD ./web/appboot /app
WORKDIR /app
RUN npm install && npm run build

FROM alpine:3.15

# https://pkgs.alpinelinux.org/packages
RUN apk add  --no-cache git bash openssh nginx

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