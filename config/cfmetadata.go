// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package config

// CFMetaData stores column family meta data
type CFMetaData struct {
	TableName     string // name of table which has this column family
	CFName        string // name of column family
	ColumnType    string // standard or super
	IndexProperty string // name sorted or timestamp sorted

	NRowKey          string
	NSuperColumnMap  string // only used in super column family
	NSuperColumnKey  string // only used in super column family
	NColumnMap       string
	NColumnKey       string
	NColumnValue     string
	NColumnTimestamp string
}

// Pretty prints and describes the column family
func (c *CFMetaData) Pretty() string {
	desc := c.NColumnMap + "(" + c.NColumnKey + "," + c.NColumnValue + "," + c.NColumnTimestamp + ")"
	if c.ColumnType == "Super" {
		desc = c.NSuperColumnMap + "(" + c.NSuperColumnKey + "," + desc + ")"
	}
	desc = c.TableName + "." + c.CFName + "(" + c.NRowKey + "," + desc + ")\n"
	desc += "Column Family Type: " + c.ColumnType + "\nColumns Sorted by: " + c.IndexProperty + "\n"
	return desc
}
