package model

import (
	"time"

	"github.comlunatictiol/rest-api-with-go/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

//var events []Event = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO 
	events (name,description,location,date_time,user_id)
	 VALUES(?,?,?,?,?)`
	papareStmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer papareStmt.Close()
	result, err := papareStmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}

	return events, nil
}
func GetEventByID(Id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, Id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil

}
func (e Event) Update() error {
	query := "UPDATE events SET name = ?, description = ?, location = ?, date_time = ? WHERE id = ?"
	pstmt, err := db.DB.Prepare(query)

	if err != nil {

		return err
	}
	defer pstmt.Close()
	_, err = pstmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}
func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	pstmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer pstmt.Close()
	_, err = pstmt.Exec(e.ID)

	return err
}
