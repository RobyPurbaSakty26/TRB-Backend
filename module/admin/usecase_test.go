package admin

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	mocks "trb-backend/mocks/module/admin"
	"trb-backend/module/entity"
)

func Test_useCase_GetAllTransaction(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		offset int
		limit  int
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	err := errors.New("e")
	res := []entity.MasterAccount{{}}
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    []entity.MasterAccount
		want1   int64
		wantErr bool
	}
	var tests []testCase
	repoMock.On("TotalDataMaster").Return(int64(10), err).Once()
	tests = append(tests, testCase{
		name:    "error get total data",
		fields:  fields{repo: repoMock},
		args:    args{},
		want:    nil,
		want1:   0,
		wantErr: true,
	})
	repoMock.On("TotalDataMaster").Return(int64(10), nil).Once()
	repoMock.On("GetAllTransaction", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, err).Once()
	tests = append(tests, testCase{
		name:    "error get all data",
		fields:  fields{repo: repoMock},
		args:    args{},
		want:    nil,
		want1:   0,
		wantErr: true,
	})
	repoMock.On("TotalDataMaster").Return(int64(10), nil).Once()
	repoMock.On("GetAllTransaction", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(res, nil).Once()
	tests = append(tests, testCase{
		name:    "success",
		fields:  fields{repo: repoMock},
		args:    args{},
		want:    res,
		want1:   int64(10),
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			got, got1, err := u.GetAllTransaction(tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransaction() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAllTransaction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_useCase_AssignRole(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		roleId uint
		userId uint
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	err := errors.New("e")
	type testCase struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	var tests []testCase
	user := &entity.User{}
	repoMock.On("GetById", mock.AnythingOfType("uint")).Return(user, err).Once()
	tests = append(tests, testCase{
		name:    "error get user id",
		fields:  fields{repo: repoMock},
		args:    args{},
		wantErr: true,
	})
	repoMock.On("GetById", mock.AnythingOfType("uint")).Return(user, nil).Once()
	repoMock.On("AssignRole", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(err).Once()
	tests = append(tests, testCase{
		name:    "error assign role",
		fields:  fields{repo: repoMock},
		args:    args{},
		wantErr: true,
	})
	repoMock.On("GetById", mock.AnythingOfType("uint")).Return(user, nil).Once()
	repoMock.On("AssignRole", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(nil).Once()
	tests = append(tests, testCase{
		name:    "success",
		fields:  fields{repo: repoMock},
		args:    args{},
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			if err := u.AssignRole(tt.args.roleId, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("AssignRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_GetAllRoles(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		offset int
		limit  int
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	err := errors.New("e")
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    []entity.Role
		want1   int64
		wantErr bool
	}
	var tests []testCase
	repoMock.On("TotalDataRole").Return(int64(10), err).Once()
	tests = append(tests, testCase{
		name:    "error get total role",
		fields:  fields{repo: repoMock},
		args:    args{},
		want:    nil,
		want1:   0,
		wantErr: true,
	})
	res := []entity.Role{{}}
	repoMock.On("TotalDataRole").Return(int64(10), nil).Once()
	repoMock.On("GetAllRoles", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(res, err).Once()
	tests = append(tests, testCase{
		name:    "error get data role",
		fields:  fields{repo: repoMock},
		args:    args{},
		want:    nil,
		want1:   0,
		wantErr: true,
	})
	repoMock.On("TotalDataRole").Return(int64(10), nil).Once()
	repoMock.On("GetAllRoles", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(res, nil).Once()
	tests = append(tests, testCase{
		name:    "error get data role",
		fields:  fields{repo: repoMock},
		args:    args{},
		want:    res,
		want1:   int64(10),
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			got, got1, err := u.GetAllRoles(tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllRoles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllRoles() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAllRoles() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_useCase_CreateRole(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		req *entity.Role
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	repoMock.On("CreateRole", mock.AnythingOfType("*entity.Role")).Return(nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{repo: repoMock},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			if err := u.CreateRole(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_CreateAccess(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		access *entity.Access
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	repoMock.On("CreateAccess", mock.AnythingOfType("*entity.Access")).Return(nil).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{repo: repoMock},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			if err := u.CreateAccess(tt.args.access); (err != nil) != tt.wantErr {
				t.Errorf("CreateAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_UpdateAccess(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		req *entity.Access
		id  uint
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	repoMock.On("UpdateAccess", mock.AnythingOfType("*entity.Access"), mock.AnythingOfType("uint")).Return(nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{repo: repoMock},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			if err := u.UpdateAccess(tt.args.req, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_GetRoleById(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		id uint
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	res := &entity.Role{}
	repoMock.On("GetRoleById", mock.AnythingOfType("uint")).Return(res, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Role
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{repo: repoMock},
			args:    args{},
			want:    res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			got, err := u.GetRoleById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoleById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoleById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_GetAllAccessByRoleId(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		id uint
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	err := errors.New("e")
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    []entity.Access
		wantErr bool
	}
	var tests []testCase
	role := &entity.Role{}
	repoMock.On("GetRoleById", mock.AnythingOfType("uint")).Return(role, err).Once()
	tests = append(tests, testCase{
		name:    "error get total data",
		fields:  fields{repo: repoMock},
		args:    args{},
		want:    nil,
		wantErr: true,
	})
	access := []entity.Access{{}}
	repoMock.On("GetRoleById", mock.AnythingOfType("uint")).Return(role, nil).Once()
	repoMock.On("GetAllAccessByRoleId", mock.AnythingOfType("uint")).Return(access, nil).Once()
	tests = append(tests, testCase{
		name:    "success",
		fields:  fields{repo: repoMock},
		args:    args{},
		want:    access,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			got, err := u.GetAllAccessByRoleId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllAccessByRoleId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllAccessByRoleId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_DeleteRole(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		id uint
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	err := errors.New("e")
	type testCase struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	var tests []testCase
	role := &entity.Role{}
	repoMock.On("GetRoleById", mock.AnythingOfType("uint")).Return(role, err).Once()
	tests = append(tests, testCase{
		name:    "error get role",
		fields:  fields{repo: repoMock},
		args:    args{},
		wantErr: true,
	})
	repoMock.On("GetRoleById", mock.AnythingOfType("uint")).Return(role, nil).Once()
	repoMock.On("DeleteAccess", mock.AnythingOfType("uint")).Return(err).Once()
	tests = append(tests, testCase{
		name:    "error delete access",
		fields:  fields{repo: repoMock},
		args:    args{},
		wantErr: true,
	})
	repoMock.On("GetRoleById", mock.AnythingOfType("uint")).Return(role, nil).Once()
	repoMock.On("DeleteAccess", mock.AnythingOfType("uint")).Return(nil).Once()
	repoMock.On("DeleteRole", mock.AnythingOfType("uint")).Return(err).Once()
	tests = append(tests, testCase{
		name:    "error delete role",
		fields:  fields{repo: repoMock},
		args:    args{},
		wantErr: true,
	})
	repoMock.On("GetRoleById", mock.AnythingOfType("uint")).Return(role, nil).Once()
	repoMock.On("DeleteAccess", mock.AnythingOfType("uint")).Return(nil).Once()
	repoMock.On("DeleteRole", mock.AnythingOfType("uint")).Return(nil).Once()
	tests = append(tests, testCase{
		name:    "error delete role",
		fields:  fields{repo: repoMock},
		args:    args{},
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			if err := u.DeleteRole(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_UpdateRole(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		role *entity.Role
		id   uint
	}
	repoMock := new(mocks.AdminRepositoryInterface)
	repoMock.On("UpdateRole",
		mock.AnythingOfType("*entity.Role"), mock.AnythingOfType("uint")).Return(nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{repo: repoMock},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			if err := u.UpdateRole(tt.args.role, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
