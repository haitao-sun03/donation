# 第一阶段：编译程序
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# 第二阶段：运行程序（使用 scratch）
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
ENV ENV=test
EXPOSE 8080
CMD ["./main"]