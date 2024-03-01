package models

import "time"

type (
	Data struct {
		CallbackURL string    `json:"callback_url,omitempty"`
		Method      string    `json:"method,omitempty"` // default POST
		Params      params    `json:"params,omitempty"`
		Body        string    `json:"body,omitempty"`
		EndTime     time.Time `json:"end_time,omitempty"`
	}

	DataListPars struct {

		// pagination
		Limit  int
		Offset int
	}
)
