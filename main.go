package main

import (
        "net/http"
        "github.com/gin-gonic/gin"
)

func main() {
        r := gin.Default()

        // 프록시 설정
        r.SetTrustedProxies([]string{"192.168.1.0/24", "10.0.0.0/8"})

        // HTML 템플릿 로드
        r.LoadHTMLGlob("*.html")

        r.GET("/", func(c *gin.Context) {
                c.HTML(http.StatusOK, "index.html", nil)
        })

        // 라우팅 정의
        r.GET("/testtest001", func(c *gin.Context) {
                c.String(http.StatusOK, "testtest001 page")
        })

        r.Run(":9000")
}
