package models

import (
	"database/sql/driver"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL

}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil

}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil

}

type NullInt struct {
	Int   int
	Valid bool
}

func (ni *NullInt) Scan(value interface{}) error {
	ni.Int, ni.Valid = value.(int)
	return nil
}

func (ni NullInt) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int, nil
}

type NullInt8 struct {
	Int8  int8
	Valid bool
}

func (ni *NullInt8) Scan(value interface{}) error {
	ni.Int8, ni.Valid = value.(int8)
	return nil
}

func (ni NullInt8) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int8, nil
}

type NullUint8 struct {
	Uint8 uint8
	Valid bool
}

func (ni *NullUint8) Scan(value interface{}) error {
	ni.Uint8, ni.Valid = value.(uint8)
	return nil
}

func (ni NullUint8) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Uint8, nil
}
