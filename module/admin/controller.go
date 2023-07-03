package admin

import (
	"errors"
	"log"
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
	getAllUser() (*response.AllUserResponse, error)
	getAllRole() (*response.ListRoleResponse, error)
	getRoleWithAccess(id string) (*response.RoleUserResponse, error)
	updateAccessUser(req *request.UpdateAccessRequest, id string) error
	UserApprove(id uint) (*response.UserApproveResponse, error)
	deleteUser(id uint) error
	createRole(req *request.UpdateAccessRequest) error
	deleteRole(id string) error
	assignRole(req request.AssignRoleRequest, id string) error
	getAllTransaction(page, limit string) (*response.MonitoringResponse, error)
	getListAccessName() (*response.ResponseAccessName, error)
	findVirtualAccountByByDate(accNo, startDate, endDate string) (*response.ResponseTransactionVitualAccount, error)
	findGiroBydate(accNo, startDate, endDate string) (*response.ResponseTransactionGiro, error)
}

func NewAdminController(usecase UseCaseAdminInterface) ControllerAdminInterface {
	return controller{
		useCase: usecase,
	}
}

func (c controller) findGiroBydate(accNo, startDate, endDate string) (*response.ResponseTransactionGiro, error) {
	// transform request
	req := request.FillterTransactionByDate{
		AccNo:     accNo,
		StartDate: startDate,
		EndDate:   endDate,
	}

	datas, err := c.useCase.findGiroByDate(&req)
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
	datas, err := c.useCase.findVirtualAccountByDate(&req)

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
	res, err := c.useCase.getListAccess()
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

func (c controller) getAllTransaction(page, limit string) (*response.MonitoringResponse, error) {
	datas, err := c.useCase.getAllTransaction(page, limit)
	if err != nil {
		return nil, err
	}

	result := response.MonitoringResponse{
		Status: "Success",
	}
	format := "02-01-2006"
	for _, data := range datas {
		tgl := data.LastUpdate.Format(format)
		//saldoGiro, _ := c.useCase.getSaldoGiro(data.AccountNo)
		//saldoVA, _ := c.useCase.getSaldoVA(data.AccountNo)
		//totalAccVA, _ := c.useCase.getTotalAccVA(data.AccountNo)
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

	_, err = c.useCase.getById(idUser)
	if err != nil {
		return err
	}

	err = c.useCase.assignRole(roleId, id)
	if err != nil {
		return err
	}
	return nil
}

func (c controller) getAllRole() (*response.ListRoleResponse, error) {
	roles, err := c.useCase.getAllRoles()
	if err != nil {
		return nil, err
	}

	result := response.ListRoleResponse{
		Status: "Success",
	}

	for _, role := range roles {
		item := response.ItemRole{
			Id:   role.ID,
			Name: role.Name,
		}
		result.Data = append(result.Data, item)
	}
	return &result, nil
}

func (c controller) createRole(req *request.UpdateAccessRequest) error {
	role := entity.Role{
		Name: req.Role,
	}
	err := c.useCase.createRole(&role)
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
		err := c.useCase.createAccess(accessReq)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c controller) getAllUser() (*response.AllUserResponse, error) {
	users, err := c.useCase.getAllUser()
	if err != nil {
		return nil, err
	}
	result := &response.AllUserResponse{
		Status: "Success",
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
	return result, nil
}

func (c controller) getRoleWithAccess(id string) (*response.RoleUserResponse, error) {
	data, err := c.useCase.getRoleById(id)

	result := &response.RoleUserResponse{
		Status: "Success",
		Data: response.ItemRoleResponse{
			Role: data.Name,
		},
	}
	accesses, err := c.useCase.getAllAccessByRoleId(id)
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
	err = c.useCase.updateRole(role, idUint)
	if err != nil {
		return err
	}

	for _, access := range req.Data {
		accessReq := &entity.Access{
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		err := c.useCase.updateAccess(accessReq, idUint)
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

	_, err = c.useCase.getRoleById(id)
	if err != nil {
		return err
	}
	err = c.useCase.deleteAccess(idUint)
	if err != nil {
		return err
	}

	err = c.useCase.deleteRole(id)
	if err != nil {
		return err
	}
	return nil
}

func (c controller) UserApprove(id uint) (*response.UserApproveResponse, error) {

	data, err := c.useCase.getById(id)
	if err != nil {
		return nil, err
	}

	err = c.useCase.userApprove(data)

	data, _ = c.useCase.getById(id)

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
	user, err := c.useCase.getById(id)
	if err != nil {
		return err
	}

	// Hapus pengguna dari use case
	err = c.useCase.deleteUser(user.ID)
	if err != nil {
		return err
	}

	return nil
}
