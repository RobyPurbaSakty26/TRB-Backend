package user

import (
	"reflect"
	"testing"
	"trb-backend/module/entity"
)

func TestUseCase_getUserAndRole(t *testing.T) {
	type fields struct {
		repo UserRepositoryInterface
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.getUserAndRole(tt.args.id)
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
