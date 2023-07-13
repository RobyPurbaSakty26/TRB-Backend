package admin

import (
	"errors"
	"log"
	"math"
	"strconv"
	"time"
	"trb-backend/module/entity"
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

type controller struct {
	useCase UseCaseAdminInterface
}

type ControllerAdminInterface interface {
	getAllUser(page, limit string) (*response.PaginateUserResponse, error)
	getAllRole(page, limit string) (*response.PaginateRole, error)
	getRoleWithAccess(id string) (*response.RoleUserResponse, error)
	updateAccessUser(req *request.UpdateAccessRequest, id string) error
	UserApprove(id uint) (*response.UserApproveResponse, error)
	deleteUser(id uint) error
	createRole(req *request.UpdateAccessRequest) error
	deleteRole(id string) error
	assignRole(req request.AssignRoleRequest, id string) error
	getListAccessName() (*response.ResponseAccessName, error)
	findVirtualAccountByByDate(accNo, startDate, endDate string) (*response.ResponseTransactionVitualAccount, error)
	findGiroBydate(accNo, startDate, endDate string) (*response.ResponseTransactionGiro, error)
	getAllTransaction(page, limit string) (*response.PaginateMonitoring, error)
	findGiroBydatePagination(accNo, startDate, endDate string, page, limit int) (*response.ResponseTransactionGiro, error)
	findVaBydatePagination(accNo, startDate, endDate string, page, limit int) (*response.ResponseTransactionVitualAccount, error)
}

func NewAdminController(usecase UseCaseAdminInterface) ControllerAdminInterface {
	return controller{
		useCase: usecase,
	}
}

func (c controller) findGiroBydatePagination(accNo, startDate, endDate string, page, limit int) (*response.ResponseTransactionGiro, error) {
	// transform request
	offset := (page - 1) * limit
	req := request.FillterTransactionByDate{
		AccNo:     accNo,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      offset,
		Limit:     limit,
	}

	count, err := c.useCase.TotalDataTransactionGiro(&req)

	if err != nil {
		return nil, err
	}

	countInt := int(count)
	totalPage := float64(countInt) / float64(limit)

	datas, err := c.useCase.FindGiroByDatePagination(&req)
	if err != nil {
		return nil, err
	}

	res := response.ResponseTransactionGiro{
		Status:    "Success",
		TotalPage: int(math.Ceil(totalPage)),
		Limit:     limit,
		Total:     countInt,
		Page:      page,
	}
	for _, data := range datas {
		// convert time []uint8 to string and adjust response
		transactionTimeStrStr := string(data.TransactionTime)
		parsedTime, err := time.Parse("15:04:05", transactionTimeStrStr)
		if err != nil {
			log.Fatal(err)
		}
		date := data.TransactionDate
		newDate := date.Format("2006-01-02")
		Newtime := parsedTime.Format("15:04:05")

		items := response.ResponseTransactionItemsGiroGetByDate{
			ID:                uint(data.Id),
			NomorRekeningGiro: data.AccountNo,
			Currency:          data.Currency,
			TanggalTransaksi:  newDate,
			Jam:               Newtime,
			Remark:            data.Remark,
			Teller:            data.TellerId,
			Category:          data.Category,
			Amount:            data.Amount,
		}

		res.Data = append(res.Data, items)
	}
	return &res, nil

}

func (c controller) findVaBydatePagination(accNo, startDate, endDate string, page, limit int) (*response.ResponseTransactionVitualAccount, error) {
	// transform request
	offset := (page - 1) * limit
	req := request.FillterTransactionByDate{
		AccNo:     accNo,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      offset,
		Limit:     limit,
	}

	count, err := c.useCase.TotalDataTransactionVa(&req)

	if err != nil {
		return nil, err
	}

	countInt := int(count)
	totalPage := float64(countInt) / float64(limit)

	datas, err := c.useCase.FindVaByDatePagination(&req)
	if err != nil {
		return nil, err
	}

	res := response.ResponseTransactionVitualAccount{
		Status:    "Success",
		TotalPage: int(math.Ceil(totalPage)),
		Limit:     limit,
		Total:     countInt,
		Page:      page,
	}
	for _, data := range datas {
		// convert time []uint8 to string and adjust response
		transactionTimeStrStr := string(data.TransactionTime)
		parsedTime, err := time.Parse("15:04:05", transactionTimeStrStr)
		if err != nil {
			log.Fatal(err)
		}
		date := data.TransactionDate
		newDate := date.Format("2006-01-02")
		Newtime := parsedTime.Format("15:04:05")

		items := response.ResponseTransactionItemsVaGetByDate{
			ID:                          uint(data.Id),
			NomorRekeningGiro:           data.AccountNo,
			NomorRekeningVirtualAccount: data.VirtualAccountNo,
			Currency:                    data.Currency,
			TanggalTransaksi:            newDate,
			Jam:                         Newtime,
			Remark:                      data.Remark,
			Teller:                      data.TellerId,
			Category:                    data.Category,
			Credit:                      data.Credit,
		}

		res.Data = append(res.Data, items)
	}
	return &res, nil

}

func (c controller) findGiroBydate(accNo, startDate, endDate string) (*response.ResponseTransactionGiro, error) {
	// transform request
	req := request.FillterTransactionByDate{
		AccNo:     accNo,
		StartDate: startDate,
		EndDate:   endDate,
	}

	datas, err := c.useCase.FindGiroByDate(&req)
	if err != nil {
		return nil, err
	}

	res := response.ResponseTransactionGiro{
		Status: "Success",
	}
	for _, data := range datas {
		// convert time []uint8 to string and adjust response
		transactionTimeStrStr := string(data.TransactionTime)
		parsedTime, err := time.Parse("15:04:05", transactionTimeStrStr)
		if err != nil {
			log.Fatal(err)
		}
		date := data.TransactionDate
		newDate := date.Format("2006-01-02")
		Newtime := parsedTime.Format("15:04:05")

		items := response.ResponseTransactionItemsGiroGetByDate{
			ID:                uint(data.Id),
			NomorRekeningGiro: data.AccountNo,
			Currency:          data.Currency,
			TanggalTransaksi:  newDate,
			Jam:               Newtime,
			Remark:            data.Remark,
			Teller:            data.TellerId,
			Category:          data.Category,
			Amount:            data.Amount,
		}

		res.Data = append(res.Data, items)
	}
	return &res, nil

}

func (c controller) findVirtualAccountByByDate(accNo, startDate, endDate string) (*response.ResponseTransactionVitualAccount, error) {
	// transform request
	req := request.FillterTransactionByDate{
		AccNo:     accNo,
		StartDate: startDate,
		EndDate:   endDate,
	}
	// call function for get data and checking error
	datas, err := c.useCase.FindVirtualAccountByDate(&req)

	if err != nil {
		return nil, err
	}
	res := response.ResponseTransactionVitualAccount{
		Status: "Success",
	}
	for _, data := range datas {
		// convert time []uint8 to string and adjust response
		transactionTimeStrStr := string(data.TransactionTime)
		parsedTime, err := time.Parse("15:04:05", transactionTimeStrStr)
		if err != nil {
			log.Fatal(err)
		}
		date := data.TransactionDate
		newDate := date.Format("2006-01-02")
		Newtime := parsedTime.Format("15:04:05")

		items := response.ResponseTransactionItemsVaGetByDate{
			ID:                          uint(data.Id),
			NomorRekeningGiro:           data.AccountNo,
			NomorRekeningVirtualAccount: data.VirtualAccountNo,
			Currency:                    data.Currency,
			TanggalTransaksi:            newDate,
			Jam:                         Newtime,
			Remark:                      data.Remark,
			Teller:                      data.TellerId,
			Category:                    data.Category,
			Credit:                      data.Credit,
		}

		res.Data = append(res.Data, items)
	}
	return &res, nil
}

func (c controller) getListAccessName() (*response.ResponseAccessName, error) {
	res, err := c.useCase.GetListAccess()
	if err != nil {
		return nil, err
	}

	result := response.ResponseAccessName{
		Status: "Success",
	}
	for _, data := range res {
		item := response.ItemAccessName{
			Name: data,
		}
		result.Data = append(result.Data, item)
	}
	return &result, nil
}
func (c controller) getAllTransaction(page, limit string) (*response.PaginateMonitoring, error) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	count, err := c.useCase.TotalDataMaster()
	if err != nil {
		return nil, errors.New("cannot get total data master")
	}
	countInt := int(count)

	totalPage := float64(countInt) / float64(limitInt)
	result := response.PaginateMonitoring{
		Page:       pageInt,
		Limit:      limitInt,
		Total:      countInt,
		TotalPages: math.Ceil(totalPage),
	}

	offset := (pageInt - 1) * limitInt
	datas, err := c.useCase.GetAllTransaction(offset, limitInt)
	if err != nil {
		return nil, err
	}
	format := "02-01-2006"
	for _, data := range datas {
		tgl := data.LastUpdate.Format(format)
		item := response.ItemMonitoring{
			NoRekeningGiro:  data.AccountNo,
			Currency:        data.Currency,
			Tanggal:         tgl,
			PosisiSaldoGiro: data.AccountBalancePosition,
			JumlahNoVA:      data.TotalVirtualAccount,
			PosisiSaldoVA:   data.VirtualAccountBalancePosition,
			Selisih:         data.AccountBalancePosition - data.VirtualAccountBalancePosition,
		}
		result.Data = append(result.Data, item)
	}

	return &result, nil
}

func (c controller) assignRole(req request.AssignRoleRequest, id string) error {
	roleId := req.RoleId
	idUserUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("cannot parse id string to uint64")
	}
	idUser := uint(idUserUint64)

	_, err = c.useCase.GetById(idUser)
	if err != nil {
		return err
	}

	err = c.useCase.AssignRole(roleId, id)
	if err != nil {
		return err
	}
	return nil
}

func (c controller) getAllRole(page, limit string) (*response.PaginateRole, error) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	count, err := c.useCase.TotalDataRole()
	if err != nil {
		return nil, errors.New("cannot get total data master")
	}
	countInt := int(count)

	totalPage := float64(countInt) / float64(limitInt)
	result := response.PaginateRole{
		Page:       pageInt,
		Limit:      limitInt,
		Total:      countInt,
		TotalPages: math.Ceil(totalPage),
	}
	offset := (pageInt - 1) * limitInt
	roles, err := c.useCase.GetAllRoles(offset, limitInt)
	if err != nil {
		return nil, errors.New("Cannot get all data roles")
	}

	for _, role := range roles {
		item := response.ItemRole{
			Id:   role.ID,
			Name: role.Name,
		}
		idStr := strconv.FormatUint(uint64(role.ID), 10)
		itemAccess, _ := c.useCase.GetAllAccessByRoleId(idStr)
		for _, data := range itemAccess {
			temp := response.AccessItem{
				Resource: data.Resource,
				CanRead:  data.CanRead,
				CanWrite: data.CanWrite,
			}
			item.Access = append(item.Access, temp)
		}
		result.Data = append(result.Data, item)
	}
	return &result, nil
}

func (c controller) createRole(req *request.UpdateAccessRequest) error {
	role := entity.Role{
		Name: req.Role,
	}
	err := c.useCase.CreateRole(&role)
	if err != nil {
		return err
	}

	for _, access := range req.Data {
		accessReq := &entity.Access{
			RoleId:   role.ID,
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		err := c.useCase.CreateAccess(accessReq)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c controller) getAllUser(page, limit string) (*response.PaginateUserResponse, error) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	count, err := c.useCase.TotalDataUser()

	if err != nil {
		return nil, errors.New("cannot get total data master")
	}
	countInt := int(count)

	totalPage := float64(countInt) / float64(limitInt)
	result := response.PaginateUserResponse{
		Page:       pageInt,
		Limit:      limitInt,
		Total:      countInt,
		TotalPages: math.Ceil(totalPage),
	}
	offset := (pageInt - 1) * limitInt
	users, err := c.useCase.GetAllUser(offset, limitInt)
	if err != nil {
		return nil, errors.New("Cannot get all data users")
	}

	for _, user := range users {
		item := response.ItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Fullname: user.Fullname,
			Email:    user.Email,
			IsActive: user.Active,
			Role:     user.Role.Name,
			RoleId:   user.RoleId,
		}
		result.Data = append(result.Data, item)
	}
	return &result, nil
}

func (c controller) getRoleWithAccess(id string) (*response.RoleUserResponse, error) {
	data, err := c.useCase.GetRoleById(id)

	result := &response.RoleUserResponse{
		Status: "Success",
		Data: response.ItemRoleResponse{
			Role: data.Name,
		},
	}
	accesses, err := c.useCase.GetAllAccessByRoleId(id)
	if err != nil {
		return nil, err
	}
	for _, access := range accesses {
		item := response.ItemAccess{
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		result.Data.Access = append(result.Data.Access, item)
	}

	return result, nil
}

func (c controller) updateAccessUser(req *request.UpdateAccessRequest, id string) error {
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("cannot parse id string to uint64")
	}
	idUint := uint(idUint64)
	role := &entity.Role{
		Name: req.Role,
	}
	err = c.useCase.UpdateRole(role, idUint)
	if err != nil {
		return err
	}

	for _, access := range req.Data {
		accessReq := &entity.Access{
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		err := c.useCase.UpdateAccess(accessReq, idUint)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c controller) deleteRole(id string) error {
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("cannot parse id string to uint64")
	}
	idUint := uint(idUint64)

	_, err = c.useCase.GetRoleById(id)
	if err != nil {
		return err
	}
	err = c.useCase.DeleteAccess(idUint)
	if err != nil {
		return err
	}

	err = c.useCase.DeleteRole(id)
	if err != nil {
		return err
	}
	return nil
}

func (c controller) UserApprove(id uint) (*response.UserApproveResponse, error) {

	data, err := c.useCase.GetById(id)
	if err != nil {
		return nil, err
	}

	err = c.useCase.UserApprove(data)

	data, _ = c.useCase.GetById(id)

	res := &response.UserApproveResponse{
		Status: "Success",
		Data: response.UserApproveItems{
			ID:       data.ID,
			Fullname: data.Fullname,
			Username: data.Username,
			Email:    data.Email,
			IsActive: data.Active,
		},
	}
	return res, nil
}

func (c controller) deleteUser(id uint) error {
	// Cek apakah pengguna dengan ID tersebut ada dalam sistem
	user, err := c.useCase.GetById(id)
	if err != nil {
		return err
	}

	// Hapus pengguna dari use case
	err = c.useCase.DeleteUser(user.ID)
	if err != nil {
		return err
	}

	return nil
}
