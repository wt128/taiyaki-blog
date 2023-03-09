package backend

import (
	"github.com/gin-gonic/gin"
	
)
func Init() {

	router().Run(":80"); // RunTLS(":80", "./server.pem", "./server.key")
}
func Router(r *gin.Engine) {
	r.GET("/article", );

}