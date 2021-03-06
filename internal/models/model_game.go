package models

import (
	"time"
)

type Game struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	// Cell State:    * `0` - New    * `1` - Playing    * `2` - Won    * `3` - Lost
	State int32 `json:"state,omitempty"`

	Started time.Time `json:"started,omitempty"`

	Finished time.Time `json:"finished,omitempty"`

	Rows int32 `json:"rows,omitempty"`

	Cols int32 `json:"cols,omitempty"`

	Mines int32 `json:"mines,omitempty"`

	Cells []Cell `json:"cells,omitempty"`
}
