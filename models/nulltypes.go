package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

func MakeNullString(s string) NullString {
	if s == "" {
		return NullString{sql.NullString{String: s, Valid: false}}
	} else {
		return NullString{sql.NullString{String: s, Valid: true}}
	}
}

type NullString struct {
	sql.NullString
}

func (v *NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.String = *s
	} else {
		v.Valid = false
	}
	return nil
}

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

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time)
	} else {
		return json.Marshal(nil)
	}
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

func (ni NullInt) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int)
	} else {
		return json.Marshal(nil)
	}
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

func (ni NullInt8) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int8)
	} else {
		return json.Marshal(nil)
	}
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

func (ni NullUint8) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Uint8)
	} else {
		return json.Marshal(nil)
	}
}
