package model

import (
	"awesomeProject1/app/server"
	"fmt"
)

type Staff struct {
	Id   int
	Name string
}

func GetAllStaff() (staffs []Staff, err error) {
	db := server.Connect()
	rows, _ := db.Query("SELECT * FROM staff")
	defer rows.Close()

	for rows.Next() {
		var staff Staff
		if err := rows.Scan(&staff.Id, &staff.Name); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		staffs = append(staffs, staff)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return
}

func GetStaffById(id int) (staff Staff, err error) {
	db := server.Connect()
	rows, _ := db.Query("SELECT * FROM staff WHERE staff.id = ", id)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&staff.Id, &staff.Name); err != nil {
			return staff, fmt.Errorf("%v", err)
		}
	}

	if err := rows.Err(); err != nil {
		return staff, fmt.Errorf("%v", err)
	}

	return
}
