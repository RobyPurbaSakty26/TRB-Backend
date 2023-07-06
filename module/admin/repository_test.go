package admin

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"testing"
	"time"
	"trb-backend/module/entity"
	"trb-backend/test"
)

func Test_repository_TotalDataUser(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type testCase struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}
	var tests []testCase

	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT count(*) FROM `users`")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    0,
		wantErr: true,
	})

	name = "success"
	var count int64 = 20
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"count(*)"}).AddRow(count),
	)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    count,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.TotalDataUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TotalDataUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_TotalDataRole(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type testCase struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}
	var tests []testCase

	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT count(*) FROM `roles`")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    0,
		wantErr: true,
	})

	name = "success"
	var count int64 = 20
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"count(*)"}).AddRow(count),
	)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    count,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.TotalDataRole()
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalDataRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TotalDataRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_TotalDataMaster(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type testCase struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}
	var tests []testCase

	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT count(*) FROM `master_account`")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    0,
		wantErr: true,
	})

	name = "success"
	var count int64 = 20
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"count(*)"}).AddRow(count),
	)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    count,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.TotalDataMaster()
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalDataMaster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TotalDataMaster() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_getListAccess(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type testCase struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}
	var tests []testCase

	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT DISTINCT resource FROM `accesses` WHERE `accesses`.`deleted_at` IS NULL")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    nil,
		wantErr: true,
	})

	name = "success"
	res := []string{
		"monitoring", "download", "user", "role",
	}

	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"resource"}).
			AddRow("monitoring").
			AddRow("download").
			AddRow("user").
			AddRow("role"),
	)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    res,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.getListAccess()
			if (err != nil) != tt.wantErr {
				t.Errorf("getListAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getListAccess() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_getAllTransaction(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		offset int
		limit  int
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    []entity.MasterAccount
		wantErr bool
	}
	var tests []testCase

	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	a := args{
		offset: 1,
		limit:  1,
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT * FROM `master_account` LIMIT 1 OFFSET 1")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		want:    nil,
		wantErr: true,
	})

	name = "success"
	res := []entity.MasterAccount{
		{
			AccountNo:                     "1",
			Currency:                      "idr",
			LastUpdate:                    time.Now(),
			AccountBalancePosition:        10,
			TotalVirtualAccount:           10,
			VirtualAccountBalancePosition: 10,
		},
	}

	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"account_no", "currency", "last_update", "account_balance_position", "total_virtual_account", "virtual_account_balance_position"}).
			AddRow("1", "idr", time.Now(), "10", "10", "10"),
	)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		want:    res,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.getAllTransaction(tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_getAllRoles(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		offset int
		limit  int
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    []entity.Role
		wantErr bool
	}
	var tests []testCase
	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	a := args{
		offset: 1,
		limit:  1,
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT * FROM `roles` WHERE `roles`.`deleted_at` IS NULL LIMIT 1 OFFSET 1")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		want:    nil,
		wantErr: true,
	})

	name = "success"
	res := []entity.Role{
		{
			Name: "admin",
		},
	}

	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"name"}).
			AddRow("admin"),
	)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		want:    res,
		wantErr: false,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.getAllRoles(tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllRoles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllRoles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_createRole(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		req *entity.Role
	}

	type testCase struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	var tests []testCase
	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	r := &entity.Role{
		Name: "admin",
	}
	a := args{
		req: r,
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("INSERT INTO `roles` (name) values (admin)")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})

	name = "success"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"name"}).
			AddRow("admin"),
	)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.createRole(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("createRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_createAccess(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		access *entity.Access
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.createAccess(tt.args.access); (err != nil) != tt.wantErr {
				t.Errorf("createAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
