FROM golang:1.19

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN go build ./cmd/bean.go

EXPOSE 8080

CMD ["./bean"]
