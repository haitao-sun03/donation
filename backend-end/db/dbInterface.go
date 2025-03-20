package db

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DBInterface 定义数据库操作接口
type DBInterface interface {
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
	WithContext(ctx context.Context) DBInterface
	Clauses(clause *clause.Clause) DBInterface
	Where(query interface{}, args ...interface{}) DBInterface
	First(out interface{}) error
	Create(value interface{}) error
	Model(value interface{}) DBInterface
	Updates(values interface{}) error
	Exec(sql string, values ...interface{}) error
}

// GormDB 实现 DBInterface
type GormDB struct {
	InnerDB *gorm.DB
}

func (g *GormDB) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return g.InnerDB.Transaction(fc, opts...)
}

func (g *GormDB) WithContext(ctx context.Context) DBInterface {
	return &GormDB{g.InnerDB.WithContext(ctx)}
}

func (g *GormDB) Clauses(clause *clause.Clause) DBInterface {
	return &GormDB{g.InnerDB.Clauses(*clause)}
}

func (g *GormDB) Where(query interface{}, args ...interface{}) DBInterface {
	return &GormDB{g.InnerDB.Where(query, args...)}
}

func (g *GormDB) First(out interface{}) error {
	return g.InnerDB.First(out).Error
}

func (g *GormDB) Create(value interface{}) error {
	return g.InnerDB.Create(value).Error
}

func (g *GormDB) Model(value interface{}) DBInterface {
	return &GormDB{g.InnerDB.Model(value)}
}

func (g *GormDB) Updates(values interface{}) error {
	return g.InnerDB.Updates(values).Error
}

func (g *GormDB) Exec(sql string, values ...interface{}) error {
	return g.InnerDB.Exec(sql, values...).Error
}
