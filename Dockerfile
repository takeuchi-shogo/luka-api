
# Golang
FROM golang:1 as golang

RUN go install github.com/cosmtrek/air@latest

# Nginx
FROM nginx:1

WORKDIR /var/www/html

COPY --from=golang /go /go
COPY --from=golang /usr/local/go /usr/local/go

COPY ./app /var/www/html/
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

COPY ./.bashrc /root/.bashrc

CMD [ "bin/dev" ]
