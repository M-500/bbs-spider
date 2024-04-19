package connpool

import (
	"context"
	"database/sql"
	"github.com/ecodeclub/ekit/syncx/atomicx"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

// @Description
// @Author 代码小学生王木木

type User struct {
	gorm.Model
	Username string
	Age      int
}
type UserV1 struct {
	gorm.Model
	Username string
	Age      int
}

func TestUseConnPool(t *testing.T) {
	sourceDB, err := gorm.Open(mysql.Open("admin:123456@tcp(192.168.1.52:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"))
	assert.NoError(t, err)
	sourceDB.AutoMigrate(&User{})
	dstDB, err := gorm.Open(mysql.Open("admin:123456@tcp(192.168.1.52:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"))
	assert.NoError(t, err)
	dstDB.AutoMigrate(&UserV1{})

	pool, err := gorm.Open(mysql.New(mysql.Config{
		Conn: &DoubleWritePool{
			src:     sourceDB.ConnPool,
			dst:     dstDB.ConnPool,
			pattern: atomicx.NewValueOf(patternSrcFirst),
		},
	}))
	if err != nil {
		panic(err)
	}
	t.Log(pool)
}

func TestDoubleWritePool_ExecContext(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DoubleWritePool{}
			got, err := d.ExecContext(tt.args.ctx, tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecContext() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleWritePool_PrepareContext(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.Stmt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DoubleWritePool{}
			got, err := d.PrepareContext(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrepareContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrepareContext() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleWritePool_QueryContext(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DoubleWritePool{}
			got, err := d.QueryContext(tt.args.ctx, tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryContext() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleWritePool_QueryRowContext(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want *sql.Row
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DoubleWritePool{}
			if got := d.QueryRowContext(tt.args.ctx, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryRowContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
