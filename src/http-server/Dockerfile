FROM golang:1.22 as builder

WORKDIR /app

COPY . .

WORKDIR /app/cmd

RUN go mod tidy && go build -o main .

FROM alpline:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/main .

EXPOSE 10607

#CMD ["./main"]

