package user

import (
	"reflect"
	"testing"
)

func TestUseCase_getAccessByRoleId(t *testing.T) {
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
		want    []entity.Access
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.getAccessByRoleId(tt.args.id)
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
