package handler

import (
	"net/http"
	"todo/pkg/helpers"
	"todo/pkg/todo_list"

	"github.com/gin-gonic/gin"
)

func (h *Handler) sighUp(c *gin.Context) {

	var input todo_list.User

	err := c.BindJSON(&input)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest , "body is empty!!!")
		return
	}

	
	_, err = h.service.Authorization.CreateUser(input)
	if err != nil{
		helpers.NewErrorResponse(c, http.StatusInternalServerError, "error while getting user id")
		return 
	}


}

func (h *Handler) sighIn(c *gin.Context) {

}
