package admin

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
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
	DownloadTransaction(c *gin.Context)
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

func (h requestAdminHandler) DownloadTransaction(c *gin.Context) {
	page := c.Query("Page")
	limit := c.Query("Limit")

	result, err := h.ctrl.getAllTransaction(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: "Failed To Retrive Data"})
		return
	}

	currentTime := time.Now()
	c.Header("Content-Disposition", fmt.Sprintf("attachment;filename=data_%s.xlsx", currentTime.Format("02-Jan-2006")))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "NoRekeningGiro")
	f.SetCellValue("Sheet1", "B1", "Currency")
	f.SetCellValue("Sheet1", "C1", "Tanggal")
	f.SetCellValue("Sheet1", "D1", "PosisiSaldoGiro")
	f.SetCellValue("Sheet1", "E1", "JumlahNoVA")
	f.SetCellValue("Sheet1", "F1", "PosisiSaldoVA")
	f.SetCellValue("Sheet1", "G1", "Selisih")

	for i, record := range result.Data {
		row := i + 2
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), record.NoRekeningGiro)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), record.Currency)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), record.Tanggal)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), record.PosisiSaldoGiro)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), record.JumlahNoVA)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), record.PosisiSaldoVA)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), record.Selisih)
	}
	err = f.Write(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: "Failed To export data to excel"})
		return
	}

	c.Status(http.StatusOK)
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

}

func (h requestAdminHandler) GetTransactionByDate(c *gin.Context) {
	from := c.Query("start_date")
	to := c.Query("end_date")
	accNo := c.Query("giro_number")
	accType := c.Query("type_account")
	page := c.Query("page")
	limit := c.Query("limit")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	page_int, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if accType != "virtual_account" {

		res, err := h.ctrl.findGiroBydatePagination(accNo, from, to, page_int, limit_int)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{Status: "Fail", Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
		return
	}

	res, err := h.ctrl.findVaBydatePagination(accNo, from, to, page_int, limit_int)
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
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
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
	page := c.Query("Page")
	limit := c.Query("Limit")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	result, err := h.ctrl.getAllRole(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: "Failed", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h requestAdminHandler) CreateRole(c *gin.Context) {
	var req request.UpdateAccessRequest

	err := c.BindJSON(&req)

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
	page := c.Query("Page")
	limit := c.Query("Limit")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	result, err := h.ctrl.getAllUser(page, limit)
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
