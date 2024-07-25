package models

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type Event struct {
	ID          int64
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Datetime    string `json:"datetime" binding:"required"`
	UserID      int64  `json:"userid"`
}

var events = []Event{}

func (e Event) Save(db *sql.DB) error {

	if db == nil {
		return errors.New("database is nill")
	}

	query := `INSERT INTO events(name,description,location,datetime,userid)
	 VALUES (?,?,?,?,?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID)

	if err != nil {
		return err
	}

	// e.ID, err = result.LastInsertId()

	// if err != nil {
	// 	return err
	// }

	//events = append(events, e)
	return nil
}

func (e Event) Update(db *sql.DB) error {

	if db == nil {
		return errors.New("database is nill")
	}

	query := `UPDATE events SET name = ? , description = ? , location = ?, datetime = ? , userid = ?  WHERE id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID, e.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete(db *sql.DB) error {

	if db == nil {
		return errors.New("database is nill")
	}

	query := `DELETE from events WHERE id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetallEvents(db *sql.DB) ([]Event, error) {

	if db == nil {
		return nil, errors.New("database is nill")
	}

	query := `SELECT * from events `

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventbyID(db *sql.DB, eventid int64) (*Event, error) {

	if db == nil {
		return nil, errors.New("database is nill")
	}

	query := `SELECT * from events where id = ?`

	row := db.QueryRow(query, eventid)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
