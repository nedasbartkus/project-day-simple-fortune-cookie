FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o backend .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/backend .

EXPOSE 9000

CMD [ "./backend" ]