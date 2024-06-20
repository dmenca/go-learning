FROM golang:1.22 as builder

WORKDIR /app

COPY . .

WORKDIR /app/cmd

# 构建Go程序，禁用CGO并将程序编译为静态二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/main .

EXPOSE 10607

#CMD ["./main"]

