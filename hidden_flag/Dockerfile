FROM golang:1.23.1

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o hidden .

CMD ["/app/hidden"]

EXPOSE 8020