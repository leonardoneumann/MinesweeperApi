package models

type PlayCell struct {
	Id int32 `json:"id"`

	// Cell State:    * `0` - Reveal    * `1` - Flag    * `2` - Question Mark
	Action int32 `json:"action"`
}
