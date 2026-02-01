package handler

import (
	"errors"
	"net/http"

	"github.com/Desiatiy10/todo-app/errs"
	"github.com/Desiatiy10/todo-app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	id, err := h.srvc.Authorization.SignUp(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	token, err := h.srvc.Authorization.SignIn(input)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, errs.ErrInvalidCredentials) {
			statusCode = http.StatusUnauthorized
		}
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}
