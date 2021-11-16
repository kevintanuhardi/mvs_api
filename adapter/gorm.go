package adapter

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Config struct {
	Master           string
	Replicas         []string
	MaxIdleConns     int
	MaxOpenConns     int
	MaxLifetimeConns int
}
type gormStarter struct {
	Opener gormOpenInterface
}
type GormStarter interface {
	ConnectToGorm(dbcfg Config, gormcfg *gorm.Config) (*gorm.DB, error)
}
type gormdb interface {
	DB() (*sql.DB, error)
	Use(plugin gorm.Plugin) error
}

func NewGormStarter() GormStarter {
	return &gormStarter{
		Opener: newGormOpener(),
	}
}

func (g *gormStarter) ConnectToGorm(dbcfg Config, gormcfg *gorm.Config) (*gorm.DB, error) {
	db, err := g.Opener.OpenConnection(dbcfg, gormcfg)
	if err != nil {
		return nil, err
	}
	// Settings connection pool
	gormdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	gormdb.SetConnMaxIdleTime(time.Duration(dbcfg.MaxIdleConns))
	gormdb.SetMaxOpenConns(dbcfg.MaxOpenConns)

	// Register Replicas
	if len(dbcfg.Replicas) > 0 {
		err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dbcfg.Master)},
			Replicas: composeReplicas(dbcfg.Replicas),
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			return nil, err
		}
	}

	return g.Opener.GetOriginal(), err
}

func composeReplicas(replicas []string) []gorm.Dialector {
	dialectors := make([]gorm.Dialector, 0)
	for _, s := range replicas {
		dialectors = append(dialectors, mysql.Open(s))
	}
	return dialectors
}
