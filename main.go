package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // 이미지가 있는 "images" 폴더를 웹에서 접근할 수 있도록 설정
    r.Static("/images", "./images")

    // 웹 페이지를 보여주는 부분
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    r.Run(":9000")
}
