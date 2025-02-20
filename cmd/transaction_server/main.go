package main

import (
	"fmt"
	"marketplace_server/cmd/transaction_server/src"
	"marketplace_server/config"
	"marketplace_server/internal/common/logs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// test
func GerData(c *gin.Context) {
	str := []byte("ok")                      // 對於[]byte感到疑惑嗎？ 因為網頁傳輸沒有string的概念，都是要轉成byte字節方式進行傳輸
	c.Data(http.StatusOK, "text/plain", str) // 指定contentType為 text/plain，就是傳輸格式為純文字啦～
}

var (
	version   string = "v0.0.1"
	buildTime string
	commitId  string
)

func main() {

	// 顯示版本
	fmt.Println("transactio-server start version:", version)
	fmt.Println("transactio-server buildTime:", buildTime)
	fmt.Println("transactio-server commitId:", commitId)

	// 載入本地 .env 檔案
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file err=", err)
	}

	// 初始化配置
	cfg := &config.Config{}
	env_flag := os.Getenv("env_flag")
	fmt.Println("env_flag=", env_flag)
	if env_flag == "1" {
		fmt.Printf("啟動env")
		cfg = config.NewEnvConfig()
	} else {
		cfg = config.NewYmlConfig("./config.yaml")
	}

	// 初始化日志
	logs.Init(cfg.Log)
	logs.Debugf("交易引擎啟動...")

	// 監聽rabbit message
	transactionEgine := src.NewTransactionEgine(cfg)
	go transactionEgine.Run()

	// 先用簡易的 gin
	gin.SetMode(cfg.Web.Mode)
	router := gin.Default()
	router.GET("/test", GerData)
	router.Run(fmt.Sprintf(":%s", cfg.Web.Port))
}
