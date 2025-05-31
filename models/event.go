package models

import (
	"github.com/daken04/Event-Booking-REST-API/db"
	"time"
)

type Event struct{
	ID int64
	Name string `Binding : "Required"`
	Description string `Binding : "Required"`
	Location string `Binding : "Required"`
	DateTime time.Time `Binding : "Required"`
	UserID int
}

var Events = []Event{}

func (e Event) Save() error {
	// Later: Save this to database
	query := `
	INSERT INTO events(Name, Description, Location, DateTime, UserID)
	VALUES (?,?,?,?,?)` // ? protects against sql injection
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	// Exec is done when we want to write on database
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetEvent(id int64) (*Event, error){
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query,id)
	var event Event
	err := row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)

	if err != nil{
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event,error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET Name = ?, Description = ?, Location = ?, DateTime = ?
	WHERE ID = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.ID)
	return err
}