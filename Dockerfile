FROM golang:1.20-alpine3.19

WORKDIR /go/src/app

COPY . .

EXPOSE 3000

RUN cat /porc/1/environ >> /go/src/app/.env

CMD [ "go", "run", "main.go" ]