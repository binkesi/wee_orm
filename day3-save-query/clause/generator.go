package clause

import (
	"fmt"
	"strings"
)

type generator func(values ...interface{}) (string, []interface{})

var generators map[Type]generator

func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[WHERE] = _where
	generators[ORDERBY] = _orderby
	generators[SELECT] = _select
	generators[LIMIT] = _limit
}

func genBindVars(nums int) string {
	var vars []string
	for i := 0; i < nums; i++ {
		vars = append(vars, "?")
	}
	res := strings.Join(vars, ", ")
	return res
}

func _insert(values ...interface{}) (string, []interface{}) {
	// INSERT INTO $tableName ($fields)
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ", ")
	return fmt.Sprintf("INSERT INTO %s, (%v)", tableName, fields), []interface{}{}
}

func _values(values ...interface{}) (string, []interface{}) {}

func _where(values ...interface{}) (string, []interface{}) {}

func _orderby(values ...interface{}) (string, []interface{}) {}

func _select(values ...interface{}) (string, []interface{}) {}

func _limit(values ...interface{}) (string, []interface{}) {}
