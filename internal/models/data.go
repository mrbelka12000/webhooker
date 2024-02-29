package models

import "time"

type Data struct {
	CallbackURL string              `json:"callback_url,omitempty"`
	Method      string              `json:"method,omitempty"` // default POST
	Params      map[string][]string `json:"params,omitempty"`
	Body        string              `json:"body,omitempty"`
	EndDate     time.Time           `json:"end_date,omitempty"`
}
