package patient

import (
	"database/sql"
)

// SQLite holds the DB connection
type SQLite struct {
	DB *sql.DB
}


func (s *SQLite) SetPatient(item Item) bool {
	stmt, _ := s.DB.Prepare("INSERT INTO patient (fullName, cpf, address, city, state, phone, email, hospital, cardNo, appointmentDate) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	stmt.Exec(item.FullName, item.Cpf, item.Address, item.City, item.State, item.Phone, item.Email, item.Hospital, item.CardNo, item.AppointmentDate)
	return true
}

// FromSQLite creates a newfeed that uses sqlite
func FromSQLite(db *sql.DB) *SQLite {
	stmt, _ := db.Prepare(`
	CREATE TABLE IF NOT EXISTS
		patient (
			ID	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			fullName	TEXT,
			cpf	TEXT,
			address	TEXT,
			city	TEXT,
			state	TEXT,
			phone	TEXT,
			email	TEXT,
			hospital	TEXT,
			cardNo	TEXT,
			appointmentDate	TEXT
		);
	`)
	stmt.Exec()
	return &SQLite{
		DB: db,
	}
}