package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/AI1411/go-psql_grpc_gql/internal/helper"
)

func addWhereEq(query *gorm.DB, columnName string, value interface{}) *gorm.DB {
	if helper.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s = ?", columnName), value)
}

func addWhereLike(query *gorm.DB, columnName string, value string) *gorm.DB {
	if helper.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s LIKE ?", columnName), "%"+value+"%")
}
