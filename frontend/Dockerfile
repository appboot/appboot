FROM node:8.12-alpine as builder

COPY ./ /app
WORKDIR /app
RUN npm install && npm run build

FROM nginx:1.17.8-alpine
RUN mkdir /app
COPY --from=builder /app/dist /app
COPY --from=builder /app/start.sh /app
COPY nginx.conf /etc/nginx/nginx.conf

WORKDIR /app

CMD sh start.sh