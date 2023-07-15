package admin

import (
	"reflect"
	"testing"
	"trb-backend/module/entity"
	"trb-backend/module/web/request"
)

func Test_useCase_getUserByEmail(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		req *request.GetByEmailUserRequset
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			got, err := u.getUserByEmail(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_TotalDataTransactionGiro(t *testing.T) {
	type fields struct {
		repo AdminRepositoryInterface
	}
	type args struct {
		req *request.FillterTransactionByDate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := useCase{
				repo: tt.fields.repo,
			}
			got, err := u.TotalDataTransactionGiro(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalDataTransactionGiro() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TotalDataTransactionGiro() got = %v, want %v", got, tt.want)
			}
		})
	}
}
