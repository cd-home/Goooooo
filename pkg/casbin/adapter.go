package casbin

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
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
	fmt.Println(args)
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

func (adapter *Adapter) SavePolicy(model model.Model) error {
	args := make([][]interface{}, 0, 64)
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			arg := _GetArgs(ptype, rule)
			args = append(args, arg)
		}
	}
	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			arg := _GetArgs(ptype, rule)
			args = append(args, arg)
		}
	}
	// first delete all
	if _, err := adapter.db.Exec(_DeleteAllPolicySQL); err != nil {
		return err
	}

	// insert new
	tx, err := adapter.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(_InsertPolicyRecordSQL)
	if err != nil {
		return err
	}
	for _, rule := range args {
		if _, err := stmt.Exec(rule...); err != nil {
			goto Rollback
		}
	}
	if err = stmt.Close(); err != nil {
		goto Rollback
	}

	if err = tx.Commit(); err != nil {
		goto Rollback
	}

Rollback:
	if err := tx.Rollback(); err != nil {
		return err
	}
	return nil
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
