FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# Build
RUN go build -o /docker-gs-ping

EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]