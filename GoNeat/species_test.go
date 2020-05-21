package GoNeat

import "testing"

func TestSpecies_GetGenomes(t *testing.T) {
	testSpecies := &Species{genomes: []*Genome{{}}}
	if len(testSpecies.GetGenomes()) != len(testSpecies.genomes) {
		t.Fatalf("Get genomes returned slice of length %v, but species genomes is of length %v",
			len(testSpecies.GetGenomes()), len(testSpecies.genomes))
	}
}

func TestSpecies_AddToGenomes(t *testing.T) {
	testSpecies := &Species{}
	testSpecies.AddToGenomes(&Genome{})
	if len(testSpecies.GetGenomes()) != 1 {
		t.Fatalf("Get genomes returned slice of length %v, but species genomes is of length 1",
			testSpecies.GetGenomes())
	}
}

func TestSpecies_GetGeneration(t *testing.T) {
	testSpecies := &Species{generation: 1}
	if testSpecies.GetGeneration() != testSpecies.generation {
		t.Fatalf("Expected test species generation to be %v, but got %v",
			testSpecies.generation, testSpecies.GetGeneration())
	}
}

func TestSpecies_IncrementGeneration(t *testing.T) {
	testSpecies := &Species{generation: 1}
	testSpecies.IncrementGeneration()
	if testSpecies.GetGeneration() != 2 {
		t.Fatalf("Expected test species generation to be 2, but got %v",
			testSpecies.generation)
	}
}

func TestSpecies_GetStagnation(t *testing.T) {
	testSpecies := &Species{stagnation: 1}
	if testSpecies.GetStagnation() != testSpecies.stagnation {
		t.Fatalf("Expected test species stagnation to be %v, but got %v",
			testSpecies.stagnation, testSpecies.GetStagnation())
	}
}

func TestSpecies_IncrementStagnation(t *testing.T) {
	testSpecies := &Species{stagnation: 1}
	testSpecies.IncrementStagnation()
	if testSpecies.GetStagnation() != 2 {
		t.Fatalf("Expected test species stagnation to be 2, but got %v",
			testSpecies.stagnation)
	}
}

func TestSpecies_GetChampion(t *testing.T) {
	testChamp := &Genome{}
	testSpecies := &Species{champion: testChamp}
	if testSpecies.GetChampion() != testChamp {
		t.Fatalf("Expected GetChampion to return champion, but it did not.")
	}
}

func TestSpecies_SetChampion(t *testing.T) {
	testGenomes := []*Genome{
		{fitness: 5},
		{fitness: 3},
		{fitness: 1},
	}
	testSpecies := &Species{genomes: testGenomes, stagnation: 3}
	testSpecies.SetChampion()

	if testSpecies.GetChampion() == nil {
		t.Fatalf("Set champion did not set champion")
	}
	if testSpecies.GetChampion() != testGenomes[0] {
		t.Fatalf("Expected set champion to select genome with fitness %v, instead selected genome with "+
			"fitness %v", testGenomes[0].GetFitness(), testSpecies.GetChampion().GetFitness())
	}

	testSpecies.AddToGenomes(&Genome{fitness: 7})
	testSpecies.SetChampion()

	if testSpecies.GetStagnation() != 0 {
		t.Fatalf("Expected stagnation to reset to 0, but its value is %v", testSpecies.GetStagnation())
	}
}

func TestSpecies_GetInnovationCounter(t *testing.T) {
	testSpecies := &Species{innovationCounter: 5}
	if testSpecies.GetInnovationCounter() != testSpecies.innovationCounter {
		t.Fatalf("Expected innovation counter to be %v, but got %v", testSpecies.innovationCounter,
			testSpecies.GetInnovationCounter())
	}
}

func TestSpecies_SetInnovationCounter(t *testing.T) {
	testSpecies := &Species{}
	testSpecies.SetInnovationCounter(5)
	if testSpecies.GetInnovationCounter() != 5 {
		t.Fatalf("Expected innovation counter to be set to 5, but got %v", testSpecies.GetInnovationCounter())
	}
}

func TestSpecies_IncrementInnovationCounter(t *testing.T) {
	testSpecies := &Species{innovationCounter: 0}
	testSpecies.IncrementInnovationCounter()
	if testSpecies.GetInnovationCounter() != 1 {
		t.Fatalf("Expected innovation counter to be incremented to 1, but got %v",
			testSpecies.GetInnovationCounter())
	}
}

func TestSpecies_InitializeInnovation(t *testing.T) {
	testGenomeOne := &Genome{innovationCounter: 5}
	testGenomeTwo := &Genome{innovationCounter: 8}
	testSpecies := &Species{}
	testSpecies.AddToGenomes(testGenomeOne)
	testSpecies.AddToGenomes(testGenomeTwo)
	testSpecies.InitializeInnovation()
	if testSpecies.GetInnovationCounter() != 8 {
		t.Fatalf("Expected innovation counter to be 8, but got %v", testSpecies.GetInnovationCounter())
	}
}

func TestSpecies_Mutate(t *testing.T) {
	testGenomeOne := InitGenome(5, 3)
	testGenomeTwo := InitGenome(5, 3)
	testGenomeOne.SetInnovationCounter(5)
	testGenomeOne.SetFitness(3.0)
	testGenomeTwo.SetInnovationCounter(8)
	testGenomeTwo.SetFitness(1.0)
	testSpecies := &Species{}
	testSpecies.AddToGenomes(testGenomeOne)
	testSpecies.AddToGenomes(testGenomeTwo)
	testSpecies.InitializeInnovation()
	testSpecies.SetChampion()
	champion := testSpecies.GetChampion()
	testSpecies.Mutate()

	champion.FeedForward()
	testGenomeTwo.FeedForward()
	for i := range champion.GetOutputs() {
		if champion.GetOutputs()[i] != testGenomeOne.GetOutputs()[i] {
			t.Fatalf("Expected champion to not be mutated, but champion outputs (%v) did not match original "+
				"outputs (%v)", champion.GetOutputs(), testGenomeOne.GetOutputs())
		}
	}
}

func TestSpecies_OrderByFitness(t *testing.T) {
	testGenomeOne := &Genome{fitness: 3}
	testGenomeTwo := &Genome{fitness: 2}
	testGenomeThree := &Genome{fitness: 1}
	testSpecies := Species{genomes: []*Genome{testGenomeOne, testGenomeTwo, testGenomeThree}}
	testSpecies.OrderByFitness()
	for i := 1; i < len(testSpecies.GetGenomes()); i++ {
		if testSpecies.GetGenomes()[i].GetFitness() < testSpecies.GetGenomes()[i-1].GetFitness() {
			t.Fatalf("Genome at index %v had a lower fitness than genom at index %v", i, i+1)
		}
	}
}

func TestSpecies_CullTheWeak(t *testing.T) {
	testGenomeOne := InitGenome(3, 5)
	testGenomeTwo := InitGenome(3, 5)
	testGenomeThree := InitGenome(3, 5)
	testGenomeOne.SetFitness(1.0)
	testGenomeTwo.SetFitness(2.0)
	testGenomeThree.SetFitness(3.0)
	testSpecies := Species{genomes: []*Genome{testGenomeOne, testGenomeTwo, testGenomeThree}, stagnation: 0}
	testSpecies.CullTheWeak()
	if len(testSpecies.GetGenomes()) != 3 {
		t.Fatalf("Expected 3 genomes, got %v", len(testSpecies.GetGenomes()))
	}
	for i := range testSpecies.GetGenomes() {
		if testSpecies.GetGenomes()[i] == testGenomeOne {
			t.Fatalf("Expected testGenomeOne to have been culled for being weak, but it was not. Found at "+
				"index %v", i)
		}
	}
}

func TestSpecies_CullTheWeak_ShouldNotReproduceIfStagnationAbove15(t *testing.T) {
	testGenomeOne := InitGenome(3, 5)
	testGenomeTwo := InitGenome(3, 5)
	testGenomeThree := InitGenome(3, 5)
	testGenomeOne.SetFitness(1.0)
	testGenomeTwo.SetFitness(2.0)
	testGenomeThree.SetFitness(3.0)
	testSpecies := Species{genomes: []*Genome{testGenomeOne, testGenomeTwo, testGenomeThree}, stagnation: 20}
	testSpecies.CullTheWeak()
	if len(testSpecies.GetGenomes()) != 2 {
		t.Fatalf("Expected 2 genomes, got %v", len(testSpecies.GetGenomes()))
	}
}

func TestSpecies_ResetStagnation(t *testing.T) {
	testSpecies := &Species{stagnation: 9}
	testSpecies.ResetStagnation()
	if testSpecies.GetStagnation() != 0 {
		t.Fatalf("Expected stagnation to be reset to 0, but it is %v", testSpecies.GetStagnation())
	}
}
