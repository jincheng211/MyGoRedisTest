FROM golang:1.20

ADD . /app

WORKDIR  /app

COPY main.go .

RUN go build -o /app/main /app/main.go

CMD ["./main"]

