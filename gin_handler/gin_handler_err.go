package gin_handler

import "github.com/gin-gonic/gin"

func GinHandlerErr(c *gin.Context) {
	GinRespondWithError(c, 400, "Something went wrong")
}
