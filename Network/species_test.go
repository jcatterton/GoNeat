package Network

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

func TestSpecies_GetFitnessCap(t *testing.T) {
	testSpecies := &Species{fitnessCap: 5}
	if testSpecies.GetFitnessCap() != testSpecies.fitnessCap {
		t.Fatalf("Expected get fitness cap (%v) to match actual fitness cap (%v), but it does not.",
			testSpecies.GetFitnessCap(), testSpecies.fitnessCap)
	}
}

func TestSpecies_SetFitnessCap(t *testing.T) {
	testSpecies := &Species{}
	testSpecies.SetFitnessCap(5.0)
	if testSpecies.GetFitnessCap() != 5.0 {
		t.Fatalf("Expected get fitness cap (%v) to match set fitness cap (5.0), but it did not.",
			testSpecies.GetFitnessCap())
	}
}

func TestSpecies_SetChampion(t *testing.T) {
	testGenomes := []*Genome{
		{fitness: 5},
		{fitness: 3},
		{fitness: 1},
	}
	testSpecies := &Species{genomes: testGenomes, stagnation: 3, fitnessCap: 6}
	testSpecies.SetChampion()

	if testSpecies.GetChampion() == nil {
		t.Fatalf("Set champion did not set champion")
	}
	if testSpecies.GetChampion() != testGenomes[0] {
		t.Fatalf("Expected set champion to select genome with fitness %v, instead selected genome with " +
			"fitness %v", testGenomes[0].GetFitness(), testSpecies.GetChampion().GetFitness())
	}
	if testSpecies.GetFitnessCap() != 6 {
		t.Fatalf("Expected fitness cap to be unaffected, but it was changed to %v", testSpecies.GetFitnessCap())
	}
	if testSpecies.GetStagnation() != 4 {
		t.Fatalf("Expected stagnation to be incremented to 4, but it was changed to %v",
			testSpecies.GetStagnation())
	}

	testSpecies.AddToGenomes(&Genome{fitness: 7})
	testSpecies.SetChampion()

	if testSpecies.GetFitnessCap() != 7 {
		t.Fatalf("Expected fitness cap to be raised to fitness of new champion, but its value is %v",
			testSpecies.GetChampion().GetFitness())
	}
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
	if testSpecies.GetInnovationCounter() <= 8 {
		t.Fatalf("Expected mutate to increase species innovation, but it is %v",
			testSpecies.GetInnovationCounter())
	}

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
		if testSpecies.GetGenomes()[i].GetFitness() < testSpecies.GetGenomes()[i - 1].GetFitness() {
			t.Fatalf("Genome at index %v had a lower fitness than genom at index %v", i, i + 1)
		}
	}
}

func TestSpecies_CullTheWeak(t *testing.T) {
	testGenomeOne := &Genome{fitness: 1}
	testGenomeTwo := &Genome{fitness: 2}
	testGenomeThree := &Genome{fitness: 3}
	testSpecies := Species{genomes: []*Genome{testGenomeOne, testGenomeTwo, testGenomeThree}}
	testSpecies.CullTheWeak()
	if len(testSpecies.GetGenomes()) != 3 {
		t.Fatalf("Expected 3 genomes, got %v", len(testSpecies.GetGenomes()))
	}
	for i := range testSpecies.GetGenomes() {
		if testSpecies.GetGenomes()[i] == testGenomeOne {
			t.Fatalf("Expected testGenomeOne to have been culled for being weak, but it was not. Found at " +
				"index %v", i)
		}
	}
}

func TestInitSpecies(t *testing.T) {
	testSpecies := InitSpecies(3, 5, 0)
	if len(testSpecies.GetGenomes()) != 10 {
		t.Fatalf("Expected species to initalize with 10 genomes, but got %v", len(testSpecies.GetGenomes()))
	}

	for i := range testSpecies.GetGenomes() {
		if len(testSpecies.GetGenomes()[i].GetNodesWithLayer(1)) != 3 {
			t.Fatalf("Expected genomes in initialized species to have 3 inputs, but got %v",
				len(testSpecies.GetGenomes()[i].GetNodesWithLayer(1)))
		}
		if len(testSpecies.GetGenomes()[i].GetNodesWithLayer(testSpecies.GetGenomes()[i].GetLayers())) != 5 {
			t.Fatalf("Expected genomes in initialized species to have 5 outputs, but got %v",
				len(testSpecies.GetGenomes()[i].GetNodesWithLayer(testSpecies.GetGenomes()[i].GetLayers())))
		}
	}
}

func TestSpecies_ResetStagnation(t *testing.T) {
	testSpecies := &Species{stagnation: 9}
	testSpecies.ResetStagnation()
	if testSpecies.GetStagnation() != 0 {
		t.Fatalf("Expected stagnation to be reset to 0, but it is %v", testSpecies.GetStagnation())
	}
}
