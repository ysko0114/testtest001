package main

import (
        "net/http"
        "github.com/gin-gonic/gin"
)

func main() {
        r := gin.Default()

        // HTML 템플릿 로드
        r.LoadHTMLGlob("*.html") // 현재 디렉터리의 모든 HTML 파일 로드
        // or r.LoadHTMLFiles("index.html") // 특정 HTML 파일 로드

        r.GET("/", func(c *gin.Context) {
                c.HTML(http.StatusOK, "index.html", nil)
        })

        r.Run(":9000")
}
