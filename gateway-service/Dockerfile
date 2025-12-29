FROM golang:1.24.6

WORKDIR /app

# Install Air
RUN go install github.com/cosmtrek/air@v1.43.0
ENV PATH="/go/bin:$PATH"

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build -o gateway ./app

EXPOSE 4000

# Run Air instead of a built binary
CMD ["air"]

