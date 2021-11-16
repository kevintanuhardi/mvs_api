package adapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormOpen func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error)
type gormOpener struct {
	db       *gorm.DB
	gormOpen gormOpen
}

type gormOpenInterface interface {
	OpenConnection(dbcfg Config, gormcfg *gorm.Config) (gormdb, error)
	GetOriginal() *gorm.DB
}

func newGormOpener() gormOpenInterface {
	return &gormOpener{
		gormOpen: gorm.Open,
	}
}
func (g *gormOpener) GetOriginal() *gorm.DB {
	return g.db
}
func (g *gormOpener) OpenConnection(dbcfg Config, gormcfg *gorm.Config) (gormdb, error) {
	db, err := g.gormOpen(mysql.Open(dbcfg.Master), gormcfg)
	g.db = db
	return db, err
}
