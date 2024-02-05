package main

import (
	"marketplace_server/cmd/transaction_engine/src"
	"marketplace_server/config"
	"marketplace_server/internal/common/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// test
func GerData(c *gin.Context) {
	str := []byte("ok")                      // 對於[]byte感到疑惑嗎？ 因為網頁傳輸沒有string的概念，都是要轉成byte字節方式進行傳輸
	c.Data(http.StatusOK, "text/plain", str) // 指定contentType為 text/plain，就是傳輸格式為純文字啦～
}

func main() {

	// 初始化配置
	cfg := config.NewConfig("./config.yaml")

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
	router.Run(":80")
}
