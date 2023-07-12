FROM golang:1.20.1 AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod download



RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/server/main.go

FROM alpine:3 AS production
COPY --from=builder /app .
CMD ["./app"]