package api

type LocationAreaDetails struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             NameURL               `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NameURL         `json:"encounter_method"`
	VersionDetails  []VersionDetail `json:"version_details"`
}

type NameURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetail struct {
	Rate    int     `json:"rate"`
	Version NameURL `json:"version"`
}

type Name struct {
	Language NameURL `json:"language"`
	Name     string  `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        NameURL `json:"pokemon"`
	VersionDetails []struct {
		EncounterDetails []EncounterDetail `json:"encounter_details"`
		MaxChance        int               `json:"max_chance"`
		Version          NameURL           `json:"version"`
	} `json:"version_details"`
}

type EncounterDetail struct {
	Chance          int     `json:"chance"`
	ConditionValues []any   `json:"condition_values"`
	MaxLevel        int     `json:"max_level"`
	Method          NameURL `json:"method"`
	MinLevel        int     `json:"min_level"`
}
