FROM golang:1.24.6

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build -o user ./app

EXPOSE 4001

CMD ["./user"]