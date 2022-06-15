package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Staff struct {
	Id   int
	Name string
}

func GetAllStaff(db *sql.DB) (staffs []Staff, err error) {
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

func GetStaffById(db *sql.DB, id int) (staff Staff, err error) {
	rows, _ := db.Query("SELECT * FROM staff WHERE staff.id = ?", id)
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

func CreateStaff(db *sql.DB, name string) (err error) {
	query := "INSERT INTO staff(name) VALUES (?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, name)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}
