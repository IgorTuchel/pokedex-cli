package pokeapi

type Pokemons struct {
	Name 	 string `json:"name"`
	Pokemon_encounters	[]struct{
		Pokemon		struct{
			Name	string `json:"name"`
			URL		string `json:"url"`
		}
	}
}

type Pokemon struct {
	Id 				int `json:"id"`
	Name 			string `json:"name"`
	Base_experience int	`json:"base_experience"`
	Height			int	`json:"height"`
	Weight 			int `json:"weight"`
	Stats 			[]struct {
		Stat 			struct {
			Name 			string `json:"name"`
			Url 			string `json:"url"`
		}
		Effort 			int `json:"effort"`
		Base_stat 		int `json:"base_stat"`
	}
	Types 			[]struct {
		Type 			struct {
			Name 			string `json:"name"`
		}
	}
}
