FROM golang:1.20-alpine3.19

WORKDIR /go/src/app

COPY . .

EXPOSE 8080

ENTRYPOINT [ ".env" ]

CMD [ "go", "run", "main.go" ]