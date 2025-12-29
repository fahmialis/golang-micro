FROM golang:1.24.6

WORKDIR /app

# Install Air
RUN go install github.com/cosmtrek/air@v1.43.0
ENV PATH="/go/bin:$PATH"

# Copy go mod files first for better cache
COPY go.mod go.sum ./
RUN go mod tidy

# Copy source code
COPY . .

EXPOSE 4001

# Run Air instead of the binary
CMD ["air"]