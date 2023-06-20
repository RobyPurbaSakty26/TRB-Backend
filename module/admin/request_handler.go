package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"trb-backend/module/web"
)

type requestAdminHandler struct {
	ctrl ControllerAdminInterface
}

type RequestHandlerAdminInterface interface {
	GetAllUser(c *gin.Context)
	GetAccessUser(c *gin.Context)
	UpdateAccessUser(c *gin.Context)
}

func NewRequestAdminHandler(ctrl ControllerAdminInterface) RequestHandlerAdminInterface {
	return &requestAdminHandler{ctrl: ctrl}
}

func DefaultRequestAdminHandler(db *gorm.DB) RequestHandlerAdminInterface {
	return NewRequestAdminHandler(
		NewAdminController(
			NewUseCase(
				NewAdminRepository(db),
			),
		),
	)
}

func (h requestAdminHandler) GetAllUser(c *gin.Context) {
	result, err := h.ctrl.getAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) GetAccessUser(c *gin.Context) {
	id := c.Param("id")

	result, err := h.ctrl.getRoleUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, web.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) UpdateAccessUser(c *gin.Context) {
	id := c.Param("id")
	var req web.UpdateAccessRequest
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	err = h.ctrl.updateAccessUser(&req, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	result, err := h.ctrl.getRoleUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
