package user

import (
	"errors"
	"reflect"
	"regexp"
	"testing"
	"time"
	"trb-backend/module/entity"
	"trb-backend/test"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func Test_repository_getAccessByRoleId(t *testing.T) {
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
		want    []entity.Access
		wantErr bool
	}

	var tests []testCase
	mockQuery, mockDb := test.NewMockQueryDB(t)
	name := "error"
	a := args{
		id: 1,
	}

	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT * FROM `accesses` WHERE role_id = ? AND `accesses`.`deleted_at` IS NUL")
	err := errors.New("error")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    a,
		want:    nil,
		wantErr: true,
	})

	name = "Success"
	res := []entity.Access{
		{
			RoleId:   1,
			Resource: "Download",
			CanRead:  true,
			CanWrite: true},
	}
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"role_id", "resource", "can_read", "can_write"}).AddRow("1", "Download", true, true),
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
			got, err := r.getAccessByRoleId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAccessByRoleId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAccessByRoleId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_save(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user *entity.User
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

	data := &entity.User{
		Fullname:   "Joe",
		Username:   "Joe",
		Email:      "Joe@mail.com",
		Password:   "password",
		RoleId:     1,
		Active:     true,
		Role:       entity.Role{},
		InputFalse: 0,
		Model: gorm.Model{
			DeletedAt: gorm.DeletedAt{},
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}

	user := args{
		user: data,
	}

	f := fields{db: mockDb}
	query := regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`fullname`,`username`,`email`,`password`,`role_id`,`active`,`input_false`) VALUES (?,?,?,?,?,?,?,?,?,?)")
	err := errors.New("error")
	mockQuery.ExpectExec(query).WithArgs(
		data.CreatedAt, // created_at
		data.UpdatedAt, // updated_at
		nil,            // deleted_at
		data.Fullname,
		data.Username,
		data.Email,
		data.Password,
		data.RoleId,
		data.Active,
		data.InputFalse,
	).WillReturnError(err)
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    user,
		wantErr: true,
	})

	name = "success"
	mockQuery.ExpectExec(query).WithArgs(
		data.CreatedAt, // created_at
		data.UpdatedAt, // updated_at
		nil,            // deleted_at
		data.Fullname,
		data.Username,
		data.Email,
		data.Password,
		data.RoleId,
		data.Active,
		data.InputFalse,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		args:    user,
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			err := r.save(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_getByEmail(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		email string
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}

	var tests []testCase

	mockQuery, mockBd := test.NewMockQueryDB(t)

	req := args{email: "joe@gmail.com"}
	f := fields{db: mockBd}
	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")
	err := errors.New("user not found")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    "Fail",
		fields:  f,
		args:    req,
		want:    nil,
		wantErr: true,
	})

	res := &entity.User{
		Fullname: "Joe",
		Username: "joe",
		Email:    "joe@mail.com",
		RoleId:   1,
	}

	mockQuery.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"fullname", "username", "email", "role_id"}).AddRow("Joe", "joe", "joe@mail.com", 1))

	tests = append(tests, testCase{
		name:    "Success",
		fields:  f,
		args:    req,
		want:    res,
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.getByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("getByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_getByUsername(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		username string
	}

	type testsCase struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}
	var tests []testsCase

	//mocking
	mockQuery, mockDb := test.NewMockQueryDB(t)
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE username = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")
	err := errors.New("error")
	mockQuery.ExpectQuery(query).WillReturnError(err)

	tests = append(tests, testsCase{
		name:   "error",
		fields: f,
		args: args{
			username: "joe",
		},
		want:    nil,
		wantErr: true,
	})

	queryRole := regexp.QuoteMeta("SELECT * FROM `roles` WHERE `roles`.`id` = ? AND `roles`.`deleted_at` IS NULL")
	mockQuery.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"fullname", "username", "email", "password", "role_id", "active"}).AddRow("Joe", "joe", "joe@mail.com", "123", 1, false))
	mockQuery.ExpectQuery(queryRole).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Admin"))
	response := &entity.User{
		Fullname: "Joe",
		Username: "joe",
		Email:    "joe@mail.com",
		Password: "123",
		RoleId:   1,
		Role: entity.Role{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "Admin",
		},
	}

	tests = append(tests, testsCase{
		name:   "Success",
		fields: f,
		args: args{
			username: "joe",
		},
		want:    response,
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.getByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("getByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_getUserAndRole(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id uint
	}
	type tesCase struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}

	var tests []tesCase
	mockQuery, mockDb := test.NewMockQueryDB(t)
	f := fields{
		db: mockDb,
	}
	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")
	err := errors.New("err")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, tesCase{
		name:   "Fail",
		fields: f,
		args: args{
			id: 1,
		},
		want:    nil,
		wantErr: true,
	})

	response := &entity.User{
		Fullname:   "Joe",
		Username:   "joe",
		Email:      "joe@mail.com",
		Password:   "123",
		RoleId:     1,
		Active:     true,
		InputFalse: 0,
	}
	mockQuery.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"fullname", "username", "email", "password", "role_id", "active", "input_false"}).AddRow("Joe", "joe", "joe@mail.com", "123", 1, true, 0))

	tests = append(tests, tesCase{
		name:   "Success",
		fields: f,
		args: args{
			id: 1,
		},
		want:    response,
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			got, err := r.getUserAndRole(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserAndRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserAndRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_updatePassword(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user     *entity.User
		password string
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	var tests []testCase
	mockQuery, mockDb := test.NewMockQueryDB(t)
	request := args{
		user:     &entity.User{},
		password: "123",
	}
	f := fields{
		db: mockDb,
	}
	query := regexp.QuoteMeta("UPDATE `users` SET `password`=?,`updated_at`=? WHERE email = ? AND `users`.`deleted_at` IS NULL")
	err := errors.New("error")
	mockQuery.ExpectQuery(query).WillReturnError(err)

	tests = append(tests, testCase{
		name:    "Fail",
		fields:  f,
		args:    request,
		wantErr: true,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.updatePassword(tt.args.user, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("updatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_updateInputFalse(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user  *entity.User
		count int
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}

	var tests []testCase
	mockQuery, mockDb := test.NewMockQueryDB(t)
	request := args{
		user: &entity.User{
			Model: gorm.Model{
				ID:        1,
				UpdatedAt: time.Now(),
			},
			Email: "joe@mail.com",
		},
		count: 1,
	}
	f := fields{
		db: mockDb,
	}
	query := regexp.QuoteMeta("UPDATE `users` SET `input_false`=?,`updated_at`=? WHERE email = ? AND `users`.`deleted_at` IS NULL AND `id` = ?")
	err := errors.New("error")
	mockQuery.ExpectQuery(query).WillReturnError(err)
	tests = append(tests, testCase{
		name:    "Fail",
		fields:  f,
		args:    request,
		wantErr: true,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.updateInputFalse(tt.args.user, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("updateInputFalse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
