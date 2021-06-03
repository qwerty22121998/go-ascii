FROM golang:1.15 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/app cmd/main.go

FROM alpine:3.10

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]