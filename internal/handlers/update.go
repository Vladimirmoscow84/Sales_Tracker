package handlers

import (
	"net/http"
	"strconv"

	"github.com/Vladimirmoscow84/Sales_Tracker/internal/model"
	"github.com/wb-go/wbf/ginext"
)

func (r *Router) updateTransactionHandler(c *ginext.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": "invalid id"})
		return
	}
	var t model.Transaction
	err = c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": err.Error()})
		return
	}
	t.ID = id
	err = r.tCRUDer.Update(c.Request.Context(), &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ginext.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
