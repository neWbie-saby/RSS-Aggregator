package gin_handler

import "github.com/gin-gonic/gin"

func GinHandlerReadiness(c *gin.Context) {
	GinRespondWithJSON(c, 200, struct{}{})
}
