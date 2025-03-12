package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // 정적 파일 제공 설정
    r.Static("/image", "./image")

    // HTML 템플릿 로드
    r.LoadHTMLGlob("testtest001/*")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    r.Run(":9000")
}
