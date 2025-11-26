package handlers

import (
	"net/http"

	"github.com/Vladimirmoscow84/Sales_Tracker/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/wb-go/wbf/ginext"
)

func (r *Router) createTransactionHandler(c *gin.Context) {
	var t model.Transaction
	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": err.Error()})
		return
	}
	id, err := r.tCRUDer.Create(c.Request.Context(), &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ginext.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ginext.H{"id": id})
}
