package user

import (
	"fmt"
	"net/http"
	"trb-backend/helpers"
	"trb-backend/module/web/request"
	"trb-backend/module/web/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**
 * Created by Goland & VS Code.
 * User : 1. Roby Purba Sakty 			: obykao26@gmail.com
		  2. Muhammad Irfan 			: mhd.irfann00@gmail.com
   		  3. Andre Rizaldi Brillianto	: andrerizaldib@gmail.com
 * Date: Saturday, 12 Juni 2023
 * Time: 08.30 AM
 * Description: BRI-CMP-Service-Backend
 **/

type RequestHandler struct {
	ctrl ControllerUserInterface
}

type RequestHandlerInterface interface {
	Create(c *gin.Context)
	GetByEmail(c *gin.Context)
	GetByUsername(c *gin.Context)
	Login(c *gin.Context)
	UpdatePassword(c *gin.Context)
	WhoIm(c *gin.Context)
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
	var req request.UserCreateRequest

	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	res, err := h.ctrl.create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) GetByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Status: "Fail", Message: "Email not found"})
		return
	}

	res, err := h.ctrl.getByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) GetByUsername(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Status: "Fail", Message: "Username not found"})
		return
	}

	res, err := h.ctrl.getByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error(), Status: "Fail"})
		return
	}

	res, err := h.ctrl.login(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error(), Status: "Fail"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) UpdatePassword(c *gin.Context) {
	var req request.UpdatePasswordRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error(), Status: "Fail"})
		return
	}

	res, err := h.ctrl.updatePassword(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) WhoIm(c *gin.Context) {
	data, err := c.Get("data")
	if err != true {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Status: "Fail", Message: "Payload not found"})
		return
	}

	dataStruct, _ := data.(helpers.PayloadJWT)
	username := dataStruct.Username

	fmt.Println(username)

	res, r := h.ctrl.getByUsername(username)
	if r != nil {
		c.JSON(http.StatusForbidden, response.ErrorResponse{Status: "Fail", Message: "Data not found"})
		return
	}

	c.JSON(http.StatusOK, res)
}
