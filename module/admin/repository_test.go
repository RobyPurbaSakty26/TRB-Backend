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

func Test_repository_assignRole(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		roleId uint
		userId string
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
	//a := args{
	//	roleId: 1,
	//	userId: "1",
	//}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("UPDATE `users` SET `role_id`=?,`updated_at`=? WHERE id = ? AND `users`.`deleted_at` IS NULL")
	err := errors.New("e")
	mockQuery.ExpectExec(query).WithArgs(0, time.Time{}, "").WillReturnError(err)
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		//args:    a,
		wantErr: true,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.assignRole(tt.args.roleId, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("assignRole() error = %v, wantErr %v", err, tt.wantErr)
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
	r := entity.Role{
		Name: "admin",
	}
	a := args{
		req: &r,
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("INSERT INTO `roles` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES (?,?,?,?)")
	err := errors.New("e")
	mockQuery.ExpectExec(query).WithArgs(time.Time{}, time.Time{}, nil, "admin").WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
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
	type testCase struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	var tests []testCase
	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	r := &entity.Access{
		RoleId:   1,
		Resource: "user",
		CanRead:  false,
		CanWrite: false,
	}
	a := args{
		access: r,
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("INSERT INTO `accesses` (`created_at`,`updated_at`,`deleted_at`,`role_id`,`resource`,`can_read`,`can_write`) VALUES (?,?,?,?,?,?,?)")
	err := errors.New("e")
	mockQuery.ExpectExec(query).
		WithArgs(time.Time{}, time.Time{}, nil, 1, "user", false, false).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})
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

func Test_repository_getAllUser(t *testing.T) {
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
		want    []entity.User
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
	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1 OFFSET 1")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})
	name = "success"
	res := []entity.User{
		{
			Fullname: "test",
			Username: "test",
			Email:    "test@gmail.com",
			Password: "Test123@",
			RoleId:   1,
			Active:   false,
			Role: entity.Role{
				Model: gorm.Model{
					ID: 1,
				},
				Name: "test",
			},
			InputFalse: 0,
		},
	}

	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"fullname", "username", "email", "password", "role_id", "active", "input_false"}).
			AddRow("test", "test", "test@gmail.com", "Test123@", 1, false, 0),
	)
	queryRole := regexp.QuoteMeta("SELECT * FROM `roles` WHERE `roles`.`id` = ? AND `roles`.`deleted_at` IS NULL")
	mockQuery.ExpectQuery(queryRole).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).
			AddRow("1", "test"),
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
			got, err := r.getAllUser(tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_getRoleById(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id string
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    *entity.Role
		wantErr bool
	}
	var tests []testCase

	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	a := args{
		id: "1",
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT * FROM `roles` WHERE `roles`.`id` = ? AND `roles`.`deleted_at` IS NULL ORDER BY `roles`.`id` LIMIT 1")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})

	name = "success"
	res := &entity.Role{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "test",
	}

	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).
			AddRow("1", "test"),
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
			got, err := r.getRoleById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRoleById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRoleById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_getAllAccessByRoleId(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id string
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    []entity.Access
		wantErr bool
	}
	var tests []testCase
	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	a := args{
		id: "1",
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT * FROM `accesses` WHERE role_id = ? AND `accesses`.`deleted_at` IS NULL")
	err := errors.New("e")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})
	name = "success"
	res := []entity.Access{
		{
			RoleId:   1,
			Resource: "Monitoring",
			CanRead:  false,
			CanWrite: false,
		},
		{
			RoleId:   1,
			Resource: "Download",
			CanRead:  false,
			CanWrite: false,
		},
	}

	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"role_id", "resource", "can_read", "can_write"}).
			AddRow("1", "Monitoring", false, false).AddRow("1", "Download", false, false),
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
			got, err := r.getAllAccessByRoleId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllAccessByRoleId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllAccessByRoleId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_updateRole(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		role *entity.Role
		id   uint
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
		Name: "test1",
	}
	a := args{
		role: r,
		id:   1,
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("UPDATE `roles` SET `name`=?,`updated_at`=? WHERE id = ? AND `roles`.`deleted_at` IS NULL")
	err := errors.New("e")
	mockQuery.ExpectExec(query).
		WithArgs("test1", time.Time{}, 1).WillReturnError(err)

	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.updateRole(tt.args.role, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("updateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_updateAccess(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		request *entity.Access
		id      uint
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
	r := &entity.Access{
		RoleId:   1,
		Resource: "user",
		CanRead:  false,
		CanWrite: false,
	}
	a := args{
		request: r,
		id:      1,
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("UPDATE `accesses` SET `can_read`=?,`can_write`=?,`updated_at`=? WHERE `accesses`.`role_id` = ? AND `accesses`.`resource` = ? AND `accesses`.`deleted_at` IS NULL")
	err := errors.New("e")
	mockQuery.ExpectExec(query).
		WithArgs(false, false, time.Time{}, 1, "user").WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.updateAccess(tt.args.request, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("updateAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_deleteAccess(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id uint
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
	a := args{
		id: 1,
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("UPDATE `accesses` SET `deleted_at`=? WHERE role_id = ? AND `accesses`.`deleted_at` IS NULL")
	err := errors.New("e")
	mockQuery.ExpectExec(query).
		WithArgs(time.Time{}, 1).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.deleteAccess(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("deleteAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_deleteRole(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id string
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
	a := args{
		id: "1",
	}
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("UPDATE `roles` SET `deleted_at`=? WHERE `roles`.`id` = ? AND `roles`.`deleted_at` IS NULL")
	err := errors.New("e")
	mockQuery.ExpectExec(query).
		WithArgs(time.Time{}, "1").WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		wantErr: true,
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.deleteRole(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("deleteRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
