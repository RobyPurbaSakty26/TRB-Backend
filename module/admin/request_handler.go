package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"trb-backend/module/web/request"
	"trb-backend/module/web/response"
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

type requestAdminHandler struct {
	ctrl ControllerAdminInterface
}

type RequestHandlerAdminInterface interface {
	GetAllUsers(c *gin.Context)
	GetAccessUser(c *gin.Context)
	UpdateAccessUser(c *gin.Context)
	UserApprove(c *gin.Context)
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

func (h requestAdminHandler) GetAllUsers(c *gin.Context) {
	result, err := h.ctrl.getAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) GetAccessUser(c *gin.Context) {
	id := c.Param("id")

	result, err := h.ctrl.getRoleUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) UpdateAccessUser(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateAccessRequest
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	err = h.ctrl.updateAccessUser(&req, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	result, err := h.ctrl.getRoleUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) UserApprove(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: "ID not found"})
		return
	}

	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error converting '%s' to int: %s\n", id, err.Error())
		return
	}
	res, err := h.ctrl.UserApprove(num)

	c.JSON(http.StatusOK, res)
}
