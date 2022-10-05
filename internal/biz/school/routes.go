package school

import "github.com/gin-gonic/gin"

func GetRoutes() func(r *gin.Engine) {
	return func(r *gin.Engine) {
		r.GET("/schools", ListSchool)
	}
}
