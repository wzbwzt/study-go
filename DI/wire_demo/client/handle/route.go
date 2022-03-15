package handle

import (
	"wireDemo/client/model"

	"github.com/gin-gonic/gin"
)

func WebRoute(r *gin.Engine) {
	r.POST("/greet", model.Greeting)
	r.POST("/goodbye", model.Goodbye)
}
