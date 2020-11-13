package mql

import (
	"time"

	"github.com/DistAlchemist/Mongongo/db"
)

func executeSetUniqueKey(tableName, columnFamilyName, rowKey, columnKey, value string) string {
	columnFamilyColumn := columnFamilyName + ":" + columnKey
	rm := db.RowMutation{tableName, rowKey, nil}
	rm.Add(columnFamilyColumn, value, time.Now().UnixNano()/int64(time.Millisecond))
	return db.Insert(rm)
}
