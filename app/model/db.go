package model

import (
	"database/sql"
	"fmt"
)

type Field struct {
	Name     string
	Type     string
	HtmlType string
}

func GetFields(db *sql.DB, table string) ([]Field, error) {

	var fields []Field

	rows, _ := db.Query("SELECT COLUMN_NAME, DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = 'staff'")
	defer rows.Close()

	for rows.Next() {
		var field Field

		if err := rows.Scan(&field.Name, &field.Type); err != nil {
			return fields, fmt.Errorf("%v", err)
		}

		fields = append(fields, field)
	}

	if err := rows.Err(); err != nil {
		return fields, fmt.Errorf("%v", err)
	}

	return fields, nil
}
