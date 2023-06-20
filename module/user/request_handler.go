package user

import (
	"fmt"
	"net/http"
	"strconv"
	"trb-backend/module/web"

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

	GetAllUsers(c *gin.Context)

	UserApprove(c *gin.Context)

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

func (h RequestHandler) UpdatePassword(c *gin.Context) {
	var req web.UpdatePasswordRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: "Fail"})
		return
	}

	res, err := h.ctrl.updatePassword(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}


func (h *RequestHandler) GetAllUsers(c *gin.Context) {
	users, err := h.ctrl.getAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": users})

func (h RequestHandler) UserApprove(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Status: "Fail", Message: "ID not found"})
		return
	}

	// Convert string to int
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error converting '%s' to int: %s\n", id, err.Error())
		return
	}

	fmt.Println(num)

	res, err := h.ctrl.UserApprove(num)

	c.JSON(http.StatusOK, res)

}
