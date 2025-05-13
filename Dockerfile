# 使用多階段建構來減少最終鏡像大小

# 建構階段
FROM golang:1.24-alpine AS builder

# 設定工作目錄
WORKDIR /app

# 複製依賴檔案
COPY go.mod go.sum ./

# 下載依賴
RUN go mod download

# 複製源碼
COPY . .

# 編譯應用
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 運行階段
FROM alpine:latest

# 安裝 ca-certificates 以支援 HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# 從建構階段複製編譯好的執行檔
COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env.example

# 設定環境變數
ENV GIN_MODE=release
ENV PORT=8080

# 開放端口
EXPOSE 8080

# 執行應用
CMD ["./main"]
