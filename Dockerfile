FROM golang:1.20

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN go build ./cmd/bean/bean.go

EXPOSE 8080

CMD ["./bean"]
