package casbin

import "fmt"

const (
	// real path
	_RuleConfPath = "../configs/rules/casbin_rabc_rule.conf"
	// for testing
	// _RuleConfPath    = "../../configs/rules/casbin_rabc_rule.conf"
	_PolicyTableName = "casbin_rule"
	_PolicyTable     = `
	CREATE TABLE %s (
		id	   INT(11) 		PRIMARY KEY AUTO_INCREMENT,
		p_type VARCHAR(32)  DEFAULT '' NOT NULL,
		v0     VARCHAR(255) DEFAULT '' NOT NULL,
		v1     VARCHAR(255) DEFAULT '' NOT NULL,
		v2     VARCHAR(255) DEFAULT '' NOT NULL,
		v3     VARCHAR(255) DEFAULT '' NOT NULL,
		v4     VARCHAR(255) DEFAULT '' NOT NULL,
		v5     VARCHAR(255) DEFAULT '' NOT NULL
	) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;`

	_CheckTableExist = "SELECT 1 FROM %s "

	_MAXParamLength = 7

	_SelectAllPolicy = "SELECT p_type,v0,v1,v2,v3,v4,v5 FROM %s"
	_InsertPolicy    = "INSERT INTO %s (p_type,v0,v1,v2,v3,v4,v5) VALUES (?,?,?,?,?,?,?)"
	_DeletePolicy    = "DELETE FROM %s WHERE p_type = ?"
	_DeleteAllPolicy = "DELETE FROM %s"
)

var (
	_CheckTableExistSQL       = fmt.Sprintf(_CheckTableExist, _PolicyTableName)
	_PolicyTableSQL           = fmt.Sprintf(_PolicyTable, _PolicyTableName)
	_SelectAllPolicyRecordSQL = fmt.Sprintf(_SelectAllPolicy, _PolicyTableName)
	_InsertPolicyRecordSQL    = fmt.Sprintf(_InsertPolicy, _PolicyTableName)
	_DeletePolicyRecordSQL    = fmt.Sprintf(_DeletePolicy, _PolicyTableName)
	_DeleteAllPolicySQL       = fmt.Sprintf(_DeletePolicy, _PolicyTableName)
)

type _CasbinRule struct {
	PType string `db:"p_type"`
	V0    string `db:"v0"`
	V1    string `db:"v1"`
	V2    string `db:"v2"`
	V3    string `db:"v3"`
	V4    string `db:"v4"`
	V5    string `db:"v5"`
}
