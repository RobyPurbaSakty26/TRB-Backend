package user

import (
	"net/http"
	"trb-backend/module/web"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	ctrl ControllerUserInterface
}

type RequestHandlerInterface interface {
	Create(c *gin.Context)
	GetByEmail(c *gin.Context)
	GetByUsername(c *gin.Context)
	Login(c *gin.Context)
}

func NewRequestHandler(ctrl ControllerUserInterface) RequestHandlerInterface {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) RequestHandlerInterface {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

func (h RequestHandler) Create(c *gin.Context) {
	var req web.UserCreateRequest

	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	res, err := h.ctrl.create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) GetByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusNotFound, web.ErrorResponse{Status: "Fail", Message: "Email not found"})
		return
	}

	res, err := h.ctrl.getByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) GetByUsername(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusNotFound, web.ErrorResponse{Status: "Fail", Message: "Username not found"})
		return
	}

	res, err := h.ctrl.getByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Status: "Fail", Message: err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Login(c *gin.Context) {
	var req web.LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: "Fail"})
		return
	}

	res, err := h.ctrl.login(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: "Fail"})
		return
	}

	c.JSON(http.StatusOK, res)
}
