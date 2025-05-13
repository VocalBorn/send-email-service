# 電子郵件寄送服務 SendEmailService
這是一個簡單的電子郵件寄送服務，使用 Go 語言實現。因應學校Server無法使用smtp寄送電子郵件的問題。
本服務會安裝在其他可以寄送SMTP的Server上，並提供內部API給學校Server使用。

## 環境需求
- Go 1.24 或以上版本
- 可用的 SMTP 伺服器（預設使用 Gmail SMTP）

## 環境設定
### 方法一：使用 .env 檔案（建議）
1. 複製 .env.example 為 .env：
```bash
cp .env.example .env
```
2. 編輯 .env 檔案，填入您的設定：
```bash
# SMTP 設定
SMTP_HOST=smtp.gmail.com    # SMTP 伺服器地址（預設：smtp.gmail.com）
SMTP_PORT=587              # SMTP 伺服器端口（預設：587）
SMTP_USERNAME=your@gmail.com # SMTP 帳號
SMTP_PASSWORD=your-password  # SMTP 密碼

# 服務設定
PORT=8080                  # API 服務端口（預設：8080）
```

### 方法二：使用系統環境變數
直接在系統中設定環境變數：
```bash
export SMTP_HOST=smtp.gmail.com
export SMTP_PORT=587
export SMTP_USERNAME=your@gmail.com
export SMTP_PASSWORD=your-password
export PORT=8080
```

## 如何執行
1. 確保已安裝 Go 1.24 或以上版本
2. 設定必要的環境變數
3. 執行服務：
```bash
go run main.go
```

## API 使用說明

### 發送郵件
- **URL**: `/send-email`
- **方法**: POST
- **請求內容格式**: JSON

**請求內容範例**：
```json
{
    "to": ["recipient1@example.com", "recipient2@example.com"],
    "subject": "測試郵件",
    "body": "<h1>Hello</h1><p>這是一封測試郵件</p>"
}
```

**回應範例**：
- 成功：
```json
{
    "message": "郵件發送成功"
}
```
- 失敗：
```json
{
    "error": "錯誤訊息"
}
```

## 注意事項
1. 使用 Gmail SMTP 時，需要在 Google 帳號設定中開啟「低安全性應用程式存取權」或使用應用程式密碼
2. 郵件內容支援 HTML 格式
3. 建議在正式環境中使用環境變數來設定 SMTP 配置
