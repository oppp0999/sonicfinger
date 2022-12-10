package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	//title값이 Home Page ("title": "Home Page" 부분)로 설정된 추가 데이터가 포함됩니다.
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})

}
