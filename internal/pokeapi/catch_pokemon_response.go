package pokeapi

type CatchPokemonResponse struct {
	BaseExperience uint64 `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Name           string `json:"name"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
