package main

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Ginのルーターを作成する
func main() {
    tracer.Start(tracer.WithServiceName("gin-sample"), tracer.WithEnv("dev"))
    defer tracer.Stop()
    router := gin.Default()

    router.Use(gintrace.Middleware("gin-sample"))

    // ルーターにHTTPハンドラーを登録する
    router.GET("/hello", helloHandler)
    router.POST("/post", postHandler)

    // ルーターを起動する
    router.Run(":8088")
}

// HTTPハンドラーを定義する
func helloHandler(c *gin.Context) {
    log.Println("hello")
    // レスポンスを書き込む
    c.String(http.StatusOK, "Hello, world!")
}

func postHandler(c *gin.Context) {
    log.Println("post-hello")
    // リクエストのボディを読み込む
    data, err := c.GetRawData()
    if err != nil {
        // エラーを返す
        c.String(http.StatusBadRequest, err.Error())
        return
    }
    // レスポンスを書き込む
    c.String(http.StatusOK, "You posted: %s", data)
}