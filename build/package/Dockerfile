# server
FROM golang:1.17.6-alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy && go build -o server cmd/server/main.go

# web
FROM node:16.14-alpine as web-builder
ADD ./web/appboot /app
WORKDIR /app
RUN npm install && npm run build

FROM alpine:3.15

# https://pkgs.alpinelinux.org/packages
RUN apk add  --no-cache git bash openssh nginx

RUN mkdir /server
RUN mkdir /app

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