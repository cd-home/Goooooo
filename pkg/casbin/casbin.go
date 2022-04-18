package casbin

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewCasbinEnforcer)

type Adapter struct {
	db *sqlx.DB
}

func NewCasbinEnforcer(db *sqlx.DB) *casbin.Enforcer {
	e, _ := casbin.NewEnforcer(_RuleConfPath, NewSqlxAdapter(db))
	return e
}

func NewSqlxAdapter(db *sqlx.DB) *Adapter {
	if db == nil {
		panic("db is nil")
	}
	if err := db.Ping(); err != nil {
		panic("db error")
	}
	if _, err := db.Queryx(_CheckTableExistSQL); err != nil {
		if _, err := db.Exec(_PolicyTableSQL); err != nil {
			log.Fatal(err)
		}
	}
	return &Adapter{db: db}
}