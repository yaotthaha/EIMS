package database

import (
	_ "embed"
	"strings"
)

//go:embed init.sql
var initSQL []byte

func parseInitSQL() []string {
	initSqlStr := string(initSQL)
	rawSqlSlice := strings.Split(initSqlStr, "\r\n")
	if len(rawSqlSlice) == 1 {
		rawSqlSlice = strings.Split(initSqlStr, "\n")
	}
	sqlSlice := make([]string, 0)
	tempSql := ""
	for k, v := range rawSqlSlice {
		v = strings.TrimSpace(v)
		if v == "" {
			if len(tempSql) > 0 {
				sqlSlice = append(sqlSlice, tempSql)
				tempSql = ""
			}
			continue
		}
		tempSql += v
		if k == len(rawSqlSlice)-1 {
			sqlSlice = append(sqlSlice, tempSql)
		}
	}
	return sqlSlice
}
