package Routes

import (
	"API_Golang/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/API_golang")
	{
		grp1.GET("employee", Controllers.GetUsers)
		grp1.POST("employee", Controllers.CreateUser)
		grp1.GET("employee/:id", Controllers.GetUserByID)
		grp1.PUT("employee/:id", Controllers.UpdateUser)
		grp1.DELETE("employee/:id", Controllers.DeleteUser)
	}
	return r
}
