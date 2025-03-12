package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // 정적 파일 제공 설정
    r.Static("/image", "./image")

    // HTML 템플릿 로드 (경로 수정)
    r.LoadHTMLGlob("testtest001/*") // templates 디렉터리에 템플릿 파일이 있는 경우

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    r.Run(":9000")
}
