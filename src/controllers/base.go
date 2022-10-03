package controllers

import (
	"Improve/src/utils"
	"github.com/gin-gonic/gin"
)

type Base struct {
}

func (b *Base) JSON(c *gin.Context, data interface{}) {
	utils.JSON(c, data)
}
func (h *Base) HandleError(c *gin.Context, err error) {
	utils.HandleError(c, err)
}

func (b *Base) Respond(c *gin.Context, v interface{}, err error) {
	if err != nil {
		b.HandleError(c, err)
	}
	b.JSON(c, v)
}
