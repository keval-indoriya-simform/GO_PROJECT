FROM golang:1.20-alpine3.19

WORKDIR /go/src/app

COPY . .

EXPOSE 8080

CMD [ "go", "run", "--env-file", "./.env", "main.go" ]