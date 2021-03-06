FROM golang:1.17.2-alpine3.14 as compiler

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM golang:1.17.2-alpine3.14

COPY --from=compiler /app/main ./app/main
ENV GO_ENV production
CMD [ "/app/main" ]
