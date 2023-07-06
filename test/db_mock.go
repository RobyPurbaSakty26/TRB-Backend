package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func NewMockQueryDB(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	conn, mockQuery, _ := sqlmock.New()

	mysqlConfig := mysql.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	}
	option := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Time{}
		},
	}

	mockDb, _ := gorm.Open(mysql.New(mysqlConfig), option)
	return mockQuery, mockDb
}
