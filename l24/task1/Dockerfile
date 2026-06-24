FROM alpine:3.23.5

RUN apk update  && \
    apk add curl nginx && \
    mkdir -p /app/static/ && \
    adduser -D -g www www

COPY ./nginx.conf /etc/nginx/nginx.conf

COPY ./app/static/index.html /app/static/index.html

RUN chown -R www:www /etc/nginx/ && \
    chown -R www:www /app


EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]