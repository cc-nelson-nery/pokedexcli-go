package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type PokeApiResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
