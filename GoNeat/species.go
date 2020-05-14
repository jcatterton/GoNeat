package GoNeat

import (
	//"math/rand"
	"math/rand"
	"sort"
)

type Species struct {
	genomes           []*Genome
	generation        int
	stagnation        int
	champion          *Genome
	innovationCounter int
}

func InitSpecies(i int, o int, g int) *Species {
	newGenomes := []*Genome{}
	for j := 0; j < g; j++ {
		newGenomes = append(newGenomes, InitGenome(i, o))
	}
	newSpecies := &Species{genomes: newGenomes, generation: g, stagnation: 0}
	return newSpecies
}

func (s *Species) GetGenomes() []*Genome {
	return s.genomes
}

func (s *Species) AddToGenomes(g *Genome) {
	s.genomes = append(s.genomes, g)
}

func (s *Species) GetGeneration() int {
	return s.generation
}

func (s *Species) IncrementGeneration() {
	s.generation = s.generation + 1
}

func (s *Species) GetStagnation() int {
	return s.stagnation
}

func (s *Species) IncrementStagnation() {
	s.stagnation = s.stagnation + 1
}

func (s *Species) ResetStagnation() {
	s.stagnation = 0
}

func (s *Species) GetChampion() *Genome {
	return s.champion
}

func (s *Species) SetChampion() {
	if s.GetChampion() == nil {
		s.champion = s.GetGenomes()[0]
	}
	s.champion.SetMutability(false)

	originalFitnessCap := s.champion.GetFitness()
	for i := range s.GetGenomes() {
		if s.GetGenomes()[i].GetFitness() > s.champion.GetFitness() {
			s.champion.SetMutability(true)
			s.champion = s.GetGenomes()[i]
			s.champion.SetMutability(false)
		}
	}

	if s.champion.GetFitness() > originalFitnessCap {
		s.ResetStagnation()
	}
}

func (s *Species) GetInnovationCounter() int {
	return s.innovationCounter
}

func (s *Species) SetInnovationCounter(i int) {
	s.innovationCounter = i
}

func (s *Species) IncrementInnovationCounter() {
	s.innovationCounter = s.innovationCounter + 1
}

func (s *Species) InitializeInnovation() {
	innovation := 0
	for g := range s.GetGenomes() {
		if s.GetGenomes()[g].GetInnovation() > innovation {
			innovation = s.GetGenomes()[g].GetInnovation()
		}
	}
	s.SetInnovationCounter(innovation)
}

func (s *Species) Mutate() {
	s.SetChampion()
	for g := range s.GetGenomes() {
		s.SetInnovationCounter(s.GetInnovationCounter() + s.GetGenomes()[g].Mutate())
		for g1 := range s.GetGenomes() {
			s.GetGenomes()[g1].SetInnovationCounter(s.GetInnovationCounter())
		}
	}
	s.IncrementStagnation()
	s.IncrementGeneration()
}

func BreedGenomes(g1 *Genome, g2 *Genome) *Genome {
	fittestParent := &Genome{}
	worstParent := &Genome{}

	if g1.GetFitness() >= g2.GetFitness() {
		fittestParent = g1.Clone()
		worstParent = g2.Clone()
	} else {
		fittestParent = g2.Clone()
		worstParent = g1.Clone()
	}
	child := &Genome{nodes: fittestParent.GetNodes(), connections: fittestParent.GetConnections()}
	for i := range worstParent.GetConnections() {
		if NodeInnovationIndex(fittestParent.GetNodes(), worstParent.GetConnections()[i].GetNodeA()) != -1 &&
			NodeInnovationIndex(fittestParent.GetNodes(), worstParent.GetConnections()[i].GetNodeB()) != -1 {
			if !fittestParent.GetNodes()[NodeInnovationIndex(child.GetNodes(), worstParent.GetConnections()[i].GetNodeA())].
				IsConnectedTo(fittestParent.GetNodes()[NodeInnovationIndex(child.GetNodes(), worstParent.GetConnections()[i].GetNodeB())]) &&
				child.GetNodes()[NodeInnovationIndex(child.GetNodes(), worstParent.GetConnections()[i].GetNodeA())].GetLayer() <
					child.GetNodes()[NodeInnovationIndex(child.GetNodes(), worstParent.GetConnections()[i].GetNodeB())].GetLayer() {
				newConnection := worstParent.GetConnections()[i].Clone()
				newConnection.SetNodeA(fittestParent.GetNodes()[NodeInnovationIndex(fittestParent.GetNodes(),
					worstParent.GetConnections()[i].GetNodeA())])
				newConnection.SetNodeB(fittestParent.GetNodes()[NodeInnovationIndex(fittestParent.GetNodes(),
					worstParent.GetConnections()[i].GetNodeB())])
				newConnection.GetNodeA().AddToOutwardConnections(newConnection)
				newConnection.GetNodeB().AddToInwardConnections(newConnection)
				child.AddConnection(newConnection)
			}
		}
	}

	child.SetLayers(fittestParent.GetLayers())
	child.SetMutability(true)
	child.SetInnovationCounter(fittestParent.GetInnovation())

	return child
}

func (s *Species) OrderByFitness() {
	sort.Slice(s.genomes, func(i, j int) bool {
		return s.genomes[i].GetFitness() < s.genomes[j].GetFitness()
	})
}

func (s *Species) CullTheWeak() {
	weaklingCounter := len(s.GetGenomes()) / 3
	s.OrderByFitness()

	s.genomes = s.genomes[weaklingCounter:]

	newGenomes := []*Genome{}
	for i := 0; i < weaklingCounter; i++ {
		newGenomes = append(newGenomes, s.BreedRandomGenomes())
	}

	s.genomes = append(s.genomes, newGenomes...)
}

func (s *Species) BreedRandomGenomes() *Genome {
	s.SetChampion()

	if len(s.genomes) == 1 {
		return s.genomes[0].Clone()
	}
	return BreedGenomes(s.genomes[rand.Intn(len(s.genomes))], s.champion)
}
