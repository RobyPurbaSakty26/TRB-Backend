package admin

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"trb-backend/helpers"
	"trb-backend/module/web/request"
	"trb-backend/module/web/response"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

type requestAdminHandler struct {
	ctrl ControllerAdminInterface
}

type RequestHandlerAdminInterface interface {
	GetAllUsers(c *gin.Context)
	GetListAccessRole(c *gin.Context)
	UpdateAccessRole(c *gin.Context)
	UserApprove(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllRoles(c *gin.Context)
	CreateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
	AssignRole(c *gin.Context)
	GetAllTransaction(c *gin.Context)
	GetListAccessName(c *gin.Context)
	//DownloadTransaction(c *gin.Context)
	GetTransactionByDate(c *gin.Context)
	DownloadTransactionByDate(c *gin.Context)
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

func (h requestAdminHandler) DownloadTransactionByDate(c *gin.Context) {
	from := c.Query("start_date")
	to := c.Query("end_date")
	accNo := c.Query("giro_number")
	accType := c.Query("type_account")

	file := excelize.NewFile()
	sheetName := "Sheet1"

	if accType != "virtual_account" {
		res, err := h.ctrl.findGiroBydate(accNo, from, to)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: err.Error()})
			return
		}
		// c.JSON(http.StatusOK, res)
		headers := helpers.GetStructTags(helpers.HeaderDownloadTransactionGiroByDate{})

		for i, header := range headers {
			_ = file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune('A'+i)), 1), header)
		}

		for i, r := range res.Data {
			rowIndex := i + 2
			v := reflect.ValueOf(r)
			for j := 0; j < v.NumField(); j++ {
				field := v.Field(j)

				err := file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune('A'+j)), rowIndex), field.Interface())
				if err != nil {
					c.JSON(http.StatusBadRequest, err)
				}
			}
		}

		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Header("Content-Disposition", "attachment; filename=transactionGiro.xlsx")

		err = file.Write(c.Writer)
		if err != nil {
			log.Println("Failed to write Excel file:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		return
	}

	res, err := h.ctrl.findVirtualAccountByByDate(accNo, from, to)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}

	headers := helpers.GetStructTags(helpers.HeaderDownloadTransactionVaByDate{})

	for i, header := range headers {
		_ = file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune('A'+i)), 1), header)
	}

	for i, r := range res.Data {
		rowIndex := i + 2
		v := reflect.ValueOf(r)
		for j := 0; j < v.NumField(); j++ {
			field := v.Field(j)

			err := file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune('A'+j)), rowIndex), field.Interface())
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=transaction.xlsx")

	err = file.Write(c.Writer)
	if err != nil {
		log.Println("Failed to write Excel file:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// c.JSON(http.StatusOK, res)

}

//func (r requestAdminHandler) DownloadTransaction(c *gin.Context) {
//
//	c.Header("Content-Disposition", "attachment;filename=data.csv")
//	c.Header("Content-Type", "text/csv")
//
//	writer := csv.NewWriter(c.Writer)
//	defer writer.Flush()
//
//	page := "5"
//	limit := "50"
//	result, err := h.ctrl.getAllTransaction(page, limit)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: "Failed To Retrive Data"})
//		return
//	}
//
//	data_header := []string{}
//
//	c.JSON(http.StatusOK, res)
//}

func (h requestAdminHandler) GetTransactionByDate(c *gin.Context) {
	from := c.Query("start_date")
	to := c.Query("end_date")
	accNo := c.Query("giro_number")
	accType := c.Query("type_account")

	if accType != "virtual_account" {

		res, err := h.ctrl.findGiroBydate(accNo, from, to)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res, err := h.ctrl.findVirtualAccountByByDate(accNo, from, to)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h requestAdminHandler) GetListAccessName(c *gin.Context) {
	res, err := h.ctrl.getListAccessName()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
func (h requestAdminHandler) GetAllTransaction(c *gin.Context) {
	page := c.Query("Page")
	limit := c.Query("Limit")
	//pageInt, _ := strconv.Atoi(page)
	//pg := (pageInt - 1) * 6
	result, err := h.ctrl.getAllTransaction(page, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) AssignRole(c *gin.Context) {
	userId := c.Param("userId")
	var req request.AssignRoleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	err = h.ctrl.assignRole(req, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "assign role success"})
}

func (h requestAdminHandler) GetAllRoles(c *gin.Context) {
	result, err := h.ctrl.getAllRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) CreateRole(c *gin.Context) {
	var req request.UpdateAccessRequest

	err := c.BindJSON(&req)
	fmt.Println(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	err = h.ctrl.createRole(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "role created"})
}

func (h requestAdminHandler) GetAllUsers(c *gin.Context) {
	result, err := h.ctrl.getAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) GetListAccessRole(c *gin.Context) {
	id := c.Param("roleId")

	result, err := h.ctrl.getRoleWithAccess(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) UpdateAccessRole(c *gin.Context) {
	id := c.Param("roleId")
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
	result, err := h.ctrl.getRoleWithAccess(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) DeleteRole(c *gin.Context) {
	id := c.Param("roleId")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: "ID not found"})
		return
	}
	err := h.ctrl.deleteRole(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Role deleted"})
}

func (h requestAdminHandler) UserApprove(c *gin.Context) {
	id := c.Param("userId")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: "ID not found"})
		return
	}

	num, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Printf("Error converting '%s' to int: %s\n", id, err.Error())
		return
	}
	idUint := uint(num)
	res, err := h.ctrl.UserApprove(idUint)

	c.JSON(http.StatusOK, res)
}

func (h requestAdminHandler) DeleteUser(c *gin.Context) {
	id := c.Param("userId")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: "ID not found"})
		return
	}

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: "Invalid ID format"})
		return
	}
	idUint := uint(userID)
	err = h.ctrl.deleteUser(idUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Fail", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted"})
}
