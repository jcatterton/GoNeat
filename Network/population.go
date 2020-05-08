package Network

import "sort"

type Population struct {
	species       []*Species
	generation    int
	grandChampion *Genome
	totalInputs   int
	totalOutputs  int
	fitnessCap    float64
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
		p.fitnessCap = p.grandChampion.GetFitness()
	}
	p.grandChampion.SetMutability(false)

	for i := range p.GetSpecies() {
		p.GetSpecies()[i].SetChampion()
	}
	for i := range p.GetSpecies() {
		if p.GetSpecies()[i].GetChampion().GetFitness() > p.fitnessCap {
			p.grandChampion.SetMutability(true)
			p.grandChampion = p.GetSpecies()[i].GetChampion()
			p.fitnessCap = p.grandChampion.GetFitness()
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

func (p *Population) GetFitnessCap() float64 {
	return p.fitnessCap
}

func (p *Population) SetFitnessCap(f float64) {
	p.fitnessCap = f
}

func (p *Population) ExtinctionEvent() {
	for i := range p.GetSpecies() {
		if p.GetSpecies()[i].GetStagnation() > 20 {
			newSpecies := &Species{stagnation: 0}
			for i := range p.GetSpecies() {
				newSpecies.AddToGenomes(p.GetSpecies()[i].BreedRandomGenomes())
				newSpecies.AddToGenomes(p.GetSpecies()[i].BreedRandomGenomes())
			}

			newSpecies.SetFitnessCap(0)

			sort.Slice(newSpecies.GetGenomes(), func(i, j int) bool {
				return newSpecies.GetGenomes()[i].GetInnovation() > newSpecies.GetGenomes()[j].GetInnovation()
			})

			newSpecies.SetInnovationCounter(newSpecies.GetGenomes()[0].GetInnovation())

			p.GetSpecies()[i] = newSpecies
		}
	}
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
