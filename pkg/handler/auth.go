package handler

import (
	"github.com/gin-gonic/gin"
	todo_list "todo"
)

func (h *Handler) sighUp(c *gin.Context) {

	var input todo_list.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

}

func (h *Handler) sighIn(c *gin.Context) {

}
