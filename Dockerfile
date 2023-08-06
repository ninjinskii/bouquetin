# Build a bqt executable, which will be available on your host root's project folder
FROM golang:1.19.0-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify
RUN chmod +x .

COPY . .

CMD ["go", "build", "-o", "./bqt"]
