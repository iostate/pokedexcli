package pokeapi

import "fmt"

type Pokemon struct {
	ID             int64
	Name           string
	BaseExperience int64 `json:"base_experience"`
	Height         int64
	Weight         int64
	Stats          []struct {
		BaseStat int64 `json:"base_stat"`
		Effort   int64
		Stat     struct {
			Name string
			URL  string
		}
	}
	Types []struct {
		Type struct {
			Name string
			URL  string
		}
	}
}

type PokemonType struct {
	Name string
	URL  string `json:"url"`
}

func (p *Pokemon) StatsMap() map[string]int64 {
	stats := make(map[string]int64)
	for _, s := range p.Stats {
		stats[s.Stat.Name] = s.BaseStat
	}
	return stats
}

func (p *Pokemon) PrintStats() {
	stats := p.StatsMap()

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", p.Name, p.Height, p.Weight)

	fmt.Printf("Stats: \n-hp: %d\n-attack: %d\n-defense: %d\n-special-attach: %d\n-special-defense: %d\n-speed: %d\n", stats["hp"], stats["attack"], stats["defense"], stats["special-attack"], stats["special-defense"], stats["speed"])
}

func (p *Pokemon) PrintTypes() {
	fmt.Printf("Types: \n")
	for _, t := range p.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
}
