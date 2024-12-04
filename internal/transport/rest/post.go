package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"posts-app/internal/domain"
	"strconv"
)

func (h *Handler) create(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input domain.Post
	if err := c.BindJSON(&input); err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Post.Create(c, userId, input)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAll(c *gin.Context) {
	posts, err := h.service.Post.GetAll(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": posts,
	})
}

func (h *Handler) getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.service.Post.GetByID(c, id)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) update(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domain.UpdatePost
	if err := c.BindJSON(&input); err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Update(c, userId, id, input); err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successfully updated",
	})
}

func (h *Handler) delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.service.Delete(c, userId, id); err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successfully deleted",
	})
}
