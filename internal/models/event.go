package models

import (
	"BookingService/internal/db"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v5"
	"strconv"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required" valid:"stringlength(0|30)"`
	Description string    `binding:"required" valid:"stringlength(0|100)"`
	Location    string    `binding:"required" valid:"stringlength(0|30)"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e *Event) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}

func (e *Event) Save() error {
	query := "INSERT INTO events (name, description, location, dateTime, userID) VALUES ($1,$2,$3,$4,$5) returning id"
	res, err := db.ConnectionPool.Query(context.Background(), query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	defer res.Close()
	ind, err := pgx.CollectOneRow(res, pgx.RowTo[int])
	if err != nil {
		return err
	}
	e.ID = ind
	return nil
}

func GetAllEvents() ([]Event, error) {
	res, err := db.ConnectionPool.Query(context.Background(), "SELECT * FROM events")
	if err != nil {
		return []Event{}, err
	}
	defer res.Close()

	resultArray, err := pgx.CollectRows(res, pgx.RowToStructByName[Event])
	if err != nil {
		return []Event{}, err
	}
	return resultArray, nil
}

func GetEventById(id int) (*Event, error) {
	var result Event
	err := db.Redis.Get("event:" + strconv.Itoa(id)).Scan(&result)

	if err != nil {
		if errors.Is(err, redis.Nil) {
			query := "SELECT * FROM events WHERE id = $1"
			rows, err := db.ConnectionPool.Query(context.Background(), query, id)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			resValue, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Event])
			if err != nil {
				return nil, err
			}
			db.Redis.SetNX("event:"+strconv.Itoa(id), resValue, time.Duration(10000))
			return &resValue, nil
		} else {
			return nil, err
		}
	}

	return &result, nil
}

func UpdateEventById(event Event) error {
	_, err := GetEventById(event.ID)
	if err != nil {
		return err
	}
	query := "UPDATE events SET name = $1, description = $2, location = $3, dateTime = $4, userID = $5 WHERE id = $6"
	rows, err := db.ConnectionPool.Query(context.Background(), query, event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

func DeleteEvent(id int) error {

	query := "DELETE FROM events WHERE id = $1"
	rows, err := db.ConnectionPool.Query(context.Background(), query, id)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

func (e *Event) Registration(userID int) error {

	query := "INSERT INTO events (eventID, userID) VALUES ($1,$2)"
	res, err := db.ConnectionPool.Query(context.Background(), query, e.ID, userID)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}

func (e *Event) DeleteRegistration(userID int) error {

	query := "DELETE FROM events WHERE eventID = $1 AND userID = $2)"
	res, err := db.ConnectionPool.Query(context.Background(), query, e.ID, userID)
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}
