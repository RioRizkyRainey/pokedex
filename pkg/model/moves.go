package model

type Moves struct {
	MoveID       int32  `json:"id"`
	MoveName     string `json:"name"`
	MovePower    *int32 `json:"power"`
	MovePp       int32  `json:"pp"`
	MoveAccuracy *int32 `json:"accuracy"`
	Type         string `json:"type"`
}
