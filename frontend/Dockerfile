FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o frontend

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/frontend .

COPY templates ./templates
COPY static ./static

EXPOSE 8080

CMD [ "./frontend" ]