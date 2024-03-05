package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type params map[string]string

func (p *params) Scan(src interface{}) (err error) {
	switch src.(type) {
	case []byte:
		return json.Unmarshal(src.([]byte), &p)
	default:
		return errors.New("incompatible type for Skills")
	}
}

func (p params) Value() (driver.Value, error) {
	return json.Marshal(p)
}
