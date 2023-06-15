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
