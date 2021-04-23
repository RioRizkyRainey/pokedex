package model

type Pokemon struct {
	PokID             int32  `json:"id"`
	PokName           string `json:"name"`
	PokHeight         int32  `json:"height"`
	PokWeight         int32  `json:"weight"`
	PokBaseExperience int32  `json:"base_experience"`
	PokAttack         int32  `json:"attack"`
	PokDefense        int32  `json:"defense"`
}
