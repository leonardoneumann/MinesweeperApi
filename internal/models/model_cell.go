package models

type Cell struct {
	Id int32 `json:"id,omitempty"`

	Row int32 `json:"row"`

	Col int32 `json:"col,omitempty"`

	NearMines int32 `json:"nearMines,omitempty"`

	// Cell State:    * `0` - Hidden    * `1` - Has Mine    * `2` - Flag    * `3` - Question Mark
	State int32 `json:"state,omitempty"`
}
