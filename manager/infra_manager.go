package manager

import (
	"database/sql"
	"fmt"
	"sync"

	"tes-mnc-bank/config"

	_ "github.com/lib/pq"
)

type InfraManager interface {
	GetDB() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg config.Config
}

var onceLoadDB sync.Once

func (i *infraManager) initDb() {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)
	onceLoadDB.Do(func() {
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			panic(err)
		}
		i.db = db
	})
	fmt.Println("DB Connected")
}
func (i *infraManager) GetDB() *sql.DB {
	return i.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initDb()
	return &infra
}
