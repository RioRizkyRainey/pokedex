package model

type Pokemon struct {
	PokID             string `json:"id"`
	PokName           string `json:"name"`
	PokHeight         int32  `json:"height"`
	PokWeight         int32  `json:"weight"`
	PokBaseExperience int32  `json:"base_experience"`
}
