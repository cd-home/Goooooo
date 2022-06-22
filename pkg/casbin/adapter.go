package casbin

import (
	"bytes"
	"strconv"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/jmoiron/sqlx"
)

func (adapter *Adapter) LoadPolicy(model model.Model) error {
	rules := make([]*_CasbinRule, 0, 128)
	if err := adapter.db.Select(&rules, _SelectAllPolicyRecordSQL); err != nil {
		return err
	}

	for _, row := range rules {
		var ruleBuffer bytes.Buffer
		ruleBuffer.Grow(64)
		ruleBuffer.WriteString(row.PType)
		args := [6]string{row.V0, row.V1, row.V2, row.V3, row.V4, row.V5}
		for _, arg := range args {
			if arg != "" {
				ruleBuffer.WriteByte(',')
				ruleBuffer.WriteString(arg)
			}
		}
		persist.LoadPolicyLine(ruleBuffer.String(), model)
	}
	return nil
}

func (adapter *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	args := _GetArgs(ptype, rule)
	_, err := adapter.db.Exec(_InsertPolicyRecordSQL, args...)
	return err
}

func (adapter *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	var sqlBuffer bytes.Buffer
	sqlBuffer.Grow(64)
	sqlBuffer.WriteString(_DeletePolicyRecordSQL)
	args := make([]interface{}, 0, 4)
	args = append(args, ptype)
	for i, arg := range rule {
		if arg != "" {
			sqlBuffer.WriteString(" AND v")
			sqlBuffer.WriteString(strconv.Itoa(i))
			sqlBuffer.WriteString(" = ?")
			args = append(args, arg)
		}
	}
	query := adapter.db.Rebind(sqlBuffer.String())
	_, err := adapter.db.Exec(query, args...)
	return err
}

func (adapter *Adapter) SavePolicy(model model.Model) (err error) {
	var tx *sqlx.Tx
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	argsP := make([]*_CasbinRule, 0, 64)
	argsG := make([]*_CasbinRule, 0, 32)
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			arg := _GetStructArg(ptype, rule)
			argsP = append(argsP, arg)
		}
	}
	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			arg := _GetStructArg(ptype, rule)
			argsG = append(argsG, arg)
		}
	}
	tx, err = adapter.db.Beginx()
	tx.MustExec(_DeleteAllPolicySQL)
	stmt, _ := tx.PrepareNamed(_InsertNamePolicySQL)
	stmt.MustExec(argsP)
	stmt.MustExec(argsG)
	return err
}

func (adapter *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	var sqlBuffer bytes.Buffer
	sqlBuffer.Grow(64)

	sqlBuffer.WriteString(_DeletePolicyRecordSQL)
	args := make([]interface{}, 0, 4)
	args = append(args, ptype)

	l := fieldIndex + len(fieldValues)

	for i := 0; i < 6; i++ {
		if fieldIndex <= i && i < l {
			v := fieldValues[i-fieldIndex]
			if v != "" {
				sqlBuffer.WriteString(" AND v")
				sqlBuffer.WriteString(strconv.Itoa(i))
				sqlBuffer.WriteString(" = ?")
				args = append(args, v)
			}
		}
	}
	query := adapter.db.Rebind(sqlBuffer.String())
	_, err := adapter.db.Exec(query, args...)
	return err
}

func _GetArgs(ptype string, rule []string) []interface{} {
	l := len(rule)
	args := make([]interface{}, _MAXParamLength)
	args[0] = ptype
	for i := 0; i < l; i++ {
		args[i+1] = rule[i]
	}
	for j := l + 1; j < _MAXParamLength; j++ {
		args[j] = ""
	}
	return args
}

func _GetStructArg(ptype string, rule []string) *_CasbinRule {
	arg := &_CasbinRule{}
	arg.PType = ptype
	arg.V0 = rule[0]
	arg.V1 = rule[1]
	switch len(rule) {
	case 3:
		arg.V2 = rule[2]
	case 4:
		arg.V2 = rule[2]
		arg.V3 = rule[3]
	case 5:
		arg.V2 = rule[2]
		arg.V3 = rule[3]
		arg.V4 = rule[4]
	case 6:
		arg.V2 = rule[2]
		arg.V3 = rule[3]
		arg.V4 = rule[4]
		arg.V5 = rule[5]
	}
	return arg
}

func _GetStructArgs(ptype string, rules [][]string) []*_CasbinRule {
	args := make([]*_CasbinRule, 0)
	for _, rule := range rules {
		args = append(args, _GetStructArg(ptype, rule))
	}
	return args
}

func (adapter *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	_rules := _GetStructArgs(ptype, rules)
	_, err := adapter.db.NamedExec(_InsertNamePolicySQL, _rules)
	return err
}

func (adapter *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) (err error) {
	var tx *sqlx.Tx
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx, err = adapter.db.Beginx()
	if err != nil {
		return
	}
	var sqlBuffer bytes.Buffer
	sqlBuffer.Grow(64)
	sqlBuffer.WriteString(_DeletePolicyRecordSQL)
	args := make([]interface{}, 0, 4)
	args = append(args, ptype)
	for _, rule := range rules {
		for i, arg := range rule {
			if arg != "" {
				sqlBuffer.WriteString(" AND v")
				sqlBuffer.WriteString(strconv.Itoa(i))
				sqlBuffer.WriteString(" = ?")
				args = append(args, arg)
			}
		}
		query := adapter.db.Rebind(sqlBuffer.String())
		tx.MustExec(query, args...)
	}
	return
}
