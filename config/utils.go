package config

import (
	"database/sql"

	"gopkg.in/mgo.v2/bson"
)

type FindResultSet struct {
	Status string   `json:"status"`
	Items  []bson.M `json:"data"`
	Note   string   `json:"note"`
}

type UpdateResult struct {
	Status string `json:"status"`
	Note   string `json:"data"`
}

func SQLToJSON(db *sql.DB, query string) FindResultSet {
	//result := bson.M{}
	rows, err := db.Query(query)
	if err != nil {
		result := FindResultSet{
			Status: "0",
			Items:  []bson.M{},
			Note:   err.Error(),
		}
		return result
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		result := FindResultSet{
			Status: "0",
			Items:  []bson.M{},
			Note:   err.Error(),
		}
		return result
	}

	count := len(columns)
	tableData := []bson.M{}
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := bson.M{}
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	result := FindResultSet{
		Status: "1",
		Items:  tableData,
		Note:   "",
	}

	return result
}
