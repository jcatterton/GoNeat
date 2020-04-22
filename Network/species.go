package Network

import (
	//"math/rand"
	"math/rand"
	"sort"
)

type Species struct {
	genomes				[]*Genome
	generation			int
	stagnation			int
	champion			*Genome
	fitnessCap			float64
	innovationCounter 	int
}

func InitSpecies(i int, o int, g int) *Species {
	newGenomes := []*Genome{}
	for j := 0; j < 10; j++ {
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
	highestFitness := 0.0
	for i := range s.GetGenomes() {
		if s.GetGenomes()[i].GetFitness() > highestFitness {
			highestFitness = s.GetGenomes()[i].GetFitness()
			s.champion = s.GetGenomes()[i]
		}
	}
	if s.GetChampion().GetFitness() > s.GetFitnessCap() {
		s.SetFitnessCap(s.GetChampion().GetFitness())
		s.ResetStagnation()
	} else {
		s.IncrementStagnation()
	}
}

func (s *Species) GetFitnessCap() float64 {
	return s.fitnessCap
}

func (s *Species) SetFitnessCap(f float64) {
	s.fitnessCap = f
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
		if s.GetGenomes()[g] != s.GetChampion() {
			s.SetInnovationCounter(s.GetInnovationCounter() + s.GetGenomes()[g].Mutate())
			for g1 := range s.GetGenomes() {
				s.GetGenomes()[g1].SetInnovationCounter(s.GetInnovationCounter())
			}
		}
	}
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
	child := &Genome{}

	//TODO: Figure out how to do this
	for i := range fittestParent.GetNodesWithLayer(1) {
		child.AddNodeWithoutIncrement(fittestParent.GetNodesWithLayer(1)[i].Clone())
	}
	for i := range fittestParent.GetNodesWithLayer(fittestParent.GetLayers()) {
		child.AddNodeWithoutIncrement(fittestParent.GetNodesWithLayer(fittestParent.GetLayers())[i].Clone())
	}
	for i := range fittestParent.GetHiddenNodes() {
		newNode := Node{}
		if NodeInnovationIndex(worstParent.GetHiddenNodes(), fittestParent.GetHiddenNodes()[i]) != -1 && i <
			len(worstParent.GetHiddenNodes()){
			if rand.Float64() >= 0.5 {
				newNode = *fittestParent.GetHiddenNodes()[i].Clone()
			} else {
				newNode = *worstParent.GetHiddenNodes()[i].Clone()
			}
		} else {
			newNode = *fittestParent.GetHiddenNodes()[i].Clone()
		}
		newNode.ClearInwardConnections()
		newNode.ClearOutwardConnections()
		child.AddNodeWithoutIncrement(&newNode)

		totalLayers := 0
		for i := range child.GetNodes() {
			if child.GetNodes()[i].GetLayer() > totalLayers {
				totalLayers = child.GetNodes()[i].GetLayer()
				child.SetLayers(totalLayers)
			}
		}
	}

	for i := range fittestParent.GetConnections() {
		if NodeInnovationIndex(child.GetNodes(), fittestParent.GetConnections()[i].GetNodeA()) != -1 &&
			NodeInnovationIndex(child.GetNodes(), fittestParent.GetConnections()[i].GetNodeB()) != -1 {
			newConnection := fittestParent.GetConnections()[i].Clone()
			newConnection.SetNodeA(child.GetNodes()[NodeInnovationIndex(child.GetNodes(),
				fittestParent.GetConnections()[i].GetNodeA())])
			newConnection.SetNodeB(child.GetNodes()[NodeInnovationIndex(child.GetNodes(),
				fittestParent.GetConnections()[i].GetNodeB())])
			newConnection.GetNodeA().AddToOutwardConnections(newConnection)
			newConnection.GetNodeB().AddToInwardConnections(newConnection)
			child.AddConnection(newConnection)
		}
	}
	for i := range worstParent.GetConnections() {
		if NodeInnovationIndex(child.GetNodes(), worstParent.GetConnections()[i].GetNodeA()) != -1 &&
			NodeInnovationIndex(child.GetNodes(), worstParent.GetConnections()[i].GetNodeB()) != -1 {
			if !child.GetNodes()[NodeInnovationIndex(child.GetNodes(), worstParent.GetConnections()[i].GetNodeA())].IsConnectedTo(
				child.GetNodes()[NodeInnovationIndex(child.GetNodes(), worstParent.GetConnections()[i].GetNodeB())]) {
				newConnection := worstParent.GetConnections()[i].Clone()
				newConnection.SetNodeA(child.GetNodes()[NodeInnovationIndex(child.GetNodes(),
					worstParent.GetConnections()[i].GetNodeA())])
				newConnection.SetNodeB(child.GetNodes()[NodeInnovationIndex(child.GetNodes(),
					worstParent.GetConnections()[i].GetNodeB())])
				newConnection.GetNodeA().AddToOutwardConnections(newConnection)
				newConnection.GetNodeB().AddToInwardConnections(newConnection)
				child.AddConnection(newConnection)
			}
		}
	}
	child.HandleDisjointNodes()
	return child
}

func (s *Species) OrderByFitness() {
	sort.Slice(s.genomes, func(i, j int) bool {
		return s.genomes[i].GetFitness() < s.genomes[j].GetFitness()
	})
}

func (s *Species) CullTheWeak() {
	weaklingCounter := len(s.GetGenomes())/3
	s.OrderByFitness()
	s.genomes = s.genomes[weaklingCounter:]

	newGenomes := []*Genome{}
	for i := 0; i < weaklingCounter; i++ {
		newGenomes = append(newGenomes, s.breedRandomGenomes())
	}

	s.genomes = append(s.genomes, newGenomes...)
}

func (s *Species) breedRandomGenomes() *Genome {
	if len(s.genomes) == 1 {
		return s.genomes[0].Clone()
	}
	genomeIndexOne := 0
	genomeIndexTwo := 0
	for genomeIndexOne == genomeIndexTwo {
		genomeIndexOne = rand.Intn(len(s.genomes))
		genomeIndexTwo = rand.Intn(len(s.genomes))
	}

	return BreedGenomes(s.genomes[genomeIndexOne], s.genomes[genomeIndexTwo])
}
