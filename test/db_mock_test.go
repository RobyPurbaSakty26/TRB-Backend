package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func TestNewMockQueryDB(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockQuery, mockDb := NewMockQueryDB(&testing.T{})
		if mockQuery == nil || mockDb == nil {
			t.Errorf("NewMockQueryDb() mockQuery = %v, mockDb %v", mockQuery, mockDb)
		}
		mockQuery.ExpectExec(regexp.QuoteMeta("INSERT")).
			WillReturnError(errors.New("e"))
		model := gorm.Model{}
		mockDb.Create(&model)
		assert.Equal(t, time.Time{}, model.CreatedAt, "model.CreatedAt")
	})
}
