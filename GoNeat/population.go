package GoNeat

import "sort"

type Population struct {
	species       []*Species
	generation    int
	grandChampion *Genome
	totalInputs   int
	totalOutputs  int
}

func InitPopulation(i int, o int) *Population {
	newSpecies := []*Species{}
	for j := 0; j < 3; j++ {
		newSpecies = append(newSpecies, InitSpecies(i, o, 0))
	}
	newPopulation := &Population{totalInputs: i, totalOutputs: o, generation: 0, species: newSpecies}
	return newPopulation
}

func (p *Population) GetSpecies() []*Species {
	return p.species
}

func (p *Population) AddToSpecies(s *Species) {
	p.species = append(p.species, s)
}

func (p *Population) GetGeneration() int {
	return p.generation
}

func (p *Population) IncrementGeneration() {
	p.generation = p.generation + 1
}

func (p *Population) GetGrandChampion() *Genome {
	return p.grandChampion
}

func (p *Population) SetGrandChampion() {
	if p.grandChampion == nil {
		p.grandChampion = p.GetAllGenomes()[0]
	}
	p.grandChampion.SetMutability(false)

	for i := range p.GetSpecies() {
		p.GetSpecies()[i].SetChampion()
	}
	for i := range p.GetSpecies() {
		if p.GetSpecies()[i].GetChampion().GetFitness() > p.grandChampion.GetFitness() {
			p.grandChampion.SetMutability(true)
			p.grandChampion = p.GetSpecies()[i].GetChampion()
			p.grandChampion.SetMutability(false)
		}
	}
}

func (p *Population) GetTotalInputs() int {
	return p.totalInputs
}

func (p *Population) GetTotalOutputs() int {
	return p.totalOutputs
}

func (p *Population) ExtinctionEvent() {
	for i := range p.GetSpecies() {
		if p.GetSpecies()[i].GetStagnation() > 20 && p.GetSpecies()[i] != p.GetChampionSpecies() {
			newSpecies := &Species{stagnation: 0}
			for i := range p.GetSpecies() {
				newSpecies.AddToGenomes(p.GetSpecies()[i].BreedRandomGenomes())
				newSpecies.AddToGenomes(p.GetSpecies()[i].BreedRandomGenomes())
			}

			sort.Slice(newSpecies.GetGenomes(), func(i, j int) bool {
				return newSpecies.GetGenomes()[i].GetInnovation() > newSpecies.GetGenomes()[j].GetInnovation()
			})

			newSpecies.SetInnovationCounter(newSpecies.GetGenomes()[0].GetInnovation())

			p.GetSpecies()[i] = newSpecies
		}
	}
}

func (p *Population) GetChampionSpecies() *Species {
	for i := range p.GetSpecies() {
		for j := range p.GetSpecies()[i].GetGenomes() {
			if p.GetSpecies()[i].GetGenomes()[j] == p.GetGrandChampion() {
				return p.GetSpecies()[i]
			}
		}
	}
	return p.GetSpecies()[0]
}

func (p *Population) Mutate() {
	for i := range p.GetSpecies() {
		p.GetSpecies()[i].Mutate()
	}
	p.IncrementGeneration()
}

func (p *Population) GetAllGenomes() []*Genome {
	allGenomes := []*Genome{}
	for i := range p.GetSpecies() {
		allGenomes = append(allGenomes, p.GetSpecies()[i].GetGenomes()...)
	}
	return allGenomes
}