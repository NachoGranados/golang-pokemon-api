package database

type Pokemon struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Type string `json:"Type"`
}

var PokemonDb = map[string]Pokemon{
	"1": Pokemon{ID: "1", Name: "Pikachu", Type: "Electric"},
	"2": Pokemon{ID: "2", Name: "Charmeleon", Type: "Fire"},
}

func PokemonDbAsValueArray() []Pokemon {
	valueArray := make([]Pokemon, len(PokemonDb))
	i := 0
	for _, value := range PokemonDb {
		valueArray[i] = value
		i++
	}
	return valueArray
}