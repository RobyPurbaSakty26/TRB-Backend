package user

import (
	"errors"
	"reflect"
	"regexp"
	"testing"
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
