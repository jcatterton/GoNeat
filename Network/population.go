package Network

type Population struct {
	species			[]*Species
	generation		int
	grandChampion	*Genome
	totalInputs		int
	totalOutputs	int
}

func InitPopulation(i int, o int) *Population {
	newSpecies := []*Species{}
	for i := 0; i < 5; i++ {
		newSpecies = append(newSpecies, InitSpecies(i, o, 0))
	}
	newPopulation := &Population{totalInputs: i, totalOutputs: i, generation: 0, species: newSpecies}
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
	fitnessCap := 0.0
	for i := range p.GetSpecies() {
		p.GetSpecies()[i].SetChampion()
		if p.GetSpecies()[i].GetChampion().GetFitness() > fitnessCap {
			p.grandChampion = p.GetSpecies()[i].GetChampion()
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
		if p.GetSpecies()[i].GetStagnation() > 10 {
			copy(p.species[i:], p.species[i + 1:])
			p.species = p.species[:len(p.species) - 1]
			p.species = append(p.species, InitSpecies(p.GetTotalInputs(), p.GetTotalOutputs(), p.GetGeneration()))
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
