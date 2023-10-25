package GoNeat

import (
	"fmt"
	"sort"
)

type Population struct {
	species       []*Species
	generation    int
	grandChampion *Genome
	totalInputs   int
	totalOutputs  int
	popCap        int
}

func InitPopulation(i int, o int, g int) *Population {
	newPopulation := &Population{totalInputs: i, totalOutputs: o, generation: 0, popCap: g}
	for k := 0; k < 4; k++ {
		newSpecies := InitSpecies(0)
		for j := 0; j < g/4; j++ {
			newSpecies.AddToGenomes(InitGenome(i, o))
		}
		newPopulation.AddToSpecies(newSpecies)
	}
	return newPopulation
}

func (p *Population) GetPopCap() int {
	return p.popCap
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
	for i := range p.GetSpecies() {
		p.GetSpecies()[i].SetChampion()
	}

	if p.grandChampion == nil {
		p.grandChampion = p.GetSpecies()[0].GetChampion()
	}

	for i := range p.GetSpecies() {
		if p.GetSpecies()[i].GetChampion().GetFitness() > p.grandChampion.GetFitness() {
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
		if len(p.GetSpecies()[i].GetGenomes()) == 0 {
			p.species = append(p.species[:i], p.species[i+1:]...)
		}
	}
}

func (p *Population) MatingSeason() {
	for i, s := range p.GetSpecies() {
		for len(s.genomes) < p.GetPopCap()/len(p.species) {
			newGenome := p.GetSpecies()[i].BreedRandomGenomes()
			newGenome.CompatibleWith(p.GetSpecies()[i].GetChampion())
			newGenome.SetMutability(true)
			p.GetSpecies()[i].AddToGenomes(newGenome)
		}
	}
	if len(p.GetAllGenomes()) < p.popCap {
		p.MatingSeason()
	}
}

func (p *Population) Speciate() {
	genomesNeedingSpeciating := []*Genome{}
	for i := range p.GetAllGenomes() {
		if p.GetAllGenomes()[i].IsMutable() {
			genomesNeedingSpeciating = append(genomesNeedingSpeciating, p.GetAllGenomes()[i])
		}
	}
	for i := range p.GetSpecies() {
		p.GetSpecies()[i].SetChampion()
		p.GetSpecies()[i].genomes = []*Genome{p.GetSpecies()[i].GetChampion()}
	}

	newSpecies := &Species{}
	for i := range genomesNeedingSpeciating {
		for j := range p.GetSpecies() {
			p.GetSpecies()[j].SetChampion()
			if genomesNeedingSpeciating[i].CompatibleWith(p.GetSpecies()[j].GetChampion()) {
				p.GetSpecies()[j].AddToGenomes(genomesNeedingSpeciating[i])
				break
			} else if !genomesNeedingSpeciating[i].CompatibleWith(p.GetSpecies()[j].GetChampion()) &&
				j == len(p.GetSpecies())-1 {
				newSpecies = InitSpecies(p.GetGeneration())
				newSpecies.AddToGenomes(genomesNeedingSpeciating[i])
				p.AddToSpecies(newSpecies)
				break
			}
		}
	}

	for i := range p.GetSpecies() {
		for j := range p.GetSpecies()[i].GetGenomes() {
			p.GetSpecies()[i].GetGenomes()[j].SetMutability(true)
		}
		p.GetSpecies()[i].SetChampion()
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

func (p *Population) NaturalSelection() {
	for i := range p.GetSpecies() {
		if p.GetSpecies()[i].GetStagnation() > 20 && p.GetSpecies()[i].GetChampion().fitness != p.GetGrandChampion().fitness {
			fmt.Println("A species has gone extinct")
			p.species[i] = InitSpecies(p.generation)
			for j := 0; j < p.GetPopCap()/4; j++ {
				p.species[i].AddToGenomes(InitGenome(p.totalInputs, p.totalOutputs))
			}
		}
		p.GetSpecies()[i].SetChampion()
		p.GetSpecies()[i].CullTheWeak(false)
	}
	p.SetGrandChampion()
	p.ExtinctionEvent()
	sort.Slice(p.species, func(i, j int) bool {
		return p.species[i].GetChampion().GetFitness() < p.species[j].GetChampion().GetFitness()
	})
	if len(p.GetSpecies()) > 10 {
		p.species = p.species[len(p.species)/3:]
	}
	p.MatingSeason()
	p.Mutate()
}
