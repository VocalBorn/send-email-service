package main

import (
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // 自動載入環境變數
)

// EmailRequest 定義發送郵件所需的請求結構
type EmailRequest struct {
	To      []string `json:"to" binding:"required"`      // 收件人郵件地址列表
	Subject string   `json:"subject" binding:"required"` // 郵件主旨
	Body    string   `json:"body" binding:"required"`    // 郵件內容（支援 HTML）
}

// SMTP 配置（建議使用環境變數設定）
var (
	smtpHost     = os.Getenv("SMTP_HOST")     // SMTP 伺服器地址
	smtpPort     = os.Getenv("SMTP_PORT")     // SMTP 伺服器端口
	smtpUsername = os.Getenv("SMTP_USERNAME") // SMTP 帳號
	smtpPassword = os.Getenv("SMTP_PASSWORD") // SMTP 密碼
)

// sendEmail 處理發送郵件的請求
func sendEmail(c *gin.Context) {
	var request EmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查 SMTP 配置
	if smtpHost == "" || smtpPort == "" || smtpUsername == "" || smtpPassword == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SMTP 配置不完整"})
		return
	}

	// 設定郵件認證資訊
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// 組合郵件標頭
	to := request.To
	subject := "Subject: " + request.Subject + "\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + request.Body)

	// 發送郵件
	mailServer := smtpHost + ":" + smtpPort
	log.Printf("正在連接到 SMTP 伺服器: %s", mailServer)
	err := smtp.SendMail(mailServer, auth, smtpUsername, to, msg)
	if err != nil {
		errorMsg := "郵件發送失敗: " + err.Error()
		log.Printf("錯誤: %s", errorMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "郵件發送成功"})
}

func main() {

	// 創建 Gin 路由器
	r := gin.Default()

	// 設定路由
	r.POST("/send-email", sendEmail)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 啟動伺服器
	log.Printf("啟動電子郵件寄送服務於 :%s Port: \n", port)
	r.Run(":" + port)
}
