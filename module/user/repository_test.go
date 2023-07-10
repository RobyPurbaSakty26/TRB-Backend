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

// func Test_repository_save(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		user *entity.User
// 	}

// 	type testCase struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}

// 	var tests []testCase
// 	mockQuery, mockDb := test.NewMockQueryDB(t)
// 	name := "error"

// 	data := &entity.User{
// 		Fullname:   "Joe",
// 		Username:   "Joe",
// 		Email:      "Joe@mail.com",
// 		Password:   "password",
// 		RoleId:     1,
// 		Active:     true,
// 		InputFalse: 0,
// 	}

// 	user := args{
// 		user: data,
// 	}

// 	f := fields{db: mockDb}
// 	query := regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`fullname`,`username`,`email`,`password`,`role_id`,`active`,`input_false`) VALUES (?,?,?,?,?,?,?,?,?,?)")
// 	err := errors.New("error")
// 	mockQuery.ExpectQuery(query).WillReturnError(err)
// 	tests = append(tests, testCase{
// 		name:    name,
// 		fields:  f,
// 		args:    user,
// 		wantErr: true,
// 	})

// 	name = "success"
// 	// mockQuery.ExpectQuery(query).WillReturnRows(
// 	// 	sqlmock.NewRows([]string{`created_at`, `updated_at`, `deleted_at`, `fullname`, `username`, `email`, `password`, `role_id`, `active`, `input_false`}).AddRow(time.Now(), time.Now(), "NULL", "Joe", "Joe", "joe@mail.com", "password", 1, true, 0),
// 	// )

// 	mockQuery.ExpectExec(query).WithArgs(
// 		mock.Anything, // created_at
// 		mock.Anything, // updated_at
// 		mock.Anything, // deleted_at
// 		data.Fullname,
// 		data.Username,
// 		data.Email,
// 		data.Password,
// 		data.RoleId,
// 		data.Active,
// 		data.InputFalse,
// 	).WillReturnResult(sqlmock.NewResult(1, 1))
// 	tests = append(tests, testCase{
// 		name:    name,
// 		fields:  f,
// 		args:    user,
// 		wantErr: false,
// 	})

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := repository{
// 				db: tt.fields.db,
// 			}
// 			if err := r.save(tt.args.user); (err != nil) != tt.wantErr {
// 				t.Errorf("save() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

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

	// Set created_at and updated_at fields to the current time
	// data.CreatedAt = time.Now()
	// data.UpdatedAt = data.CreatedAt
	// data.DeletedAt = gorm.DeletedAt{}

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
