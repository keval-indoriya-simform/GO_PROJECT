FROM golang:1.20-alpine3.19

WORKDIR /go/src/app

COPY . .

EXPOSE 3000

RUN env >> .env

RUN cat .env

CMD [ "go", "run", "main.go" ]