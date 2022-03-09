package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/wu-xie-888/micro-wire/client/model"
)

func WebRoute(r *gin.Engine) {
	r.POST("/greet", model.Greeting)
	r.POST("/goodbye", model.Goodbye)
}
