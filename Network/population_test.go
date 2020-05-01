package Network

import "testing"

func TestInitPopulation(t *testing.T) {
	testPop := InitPopulation(5, 3)
	if len(testPop.GetSpecies()) != 5 {
		t.Fatalf("Expected population to initialize with 5 species, got %v", len(testPop.GetSpecies()))
	}
}

func TestPopulation_GetSpecies(t *testing.T) {
	testSpecies := []*Species{InitSpecies(5, 3, 1)}
	testPop := &Population{species: testSpecies}
	if len(testPop.GetSpecies()) != 1 {
		t.Fatalf("Expected 1 species in test pop, got %v", len(testPop.GetSpecies()))
	}
	if testPop.GetSpecies()[0] != testSpecies[0] {
		t.Fatalf("Get species did not return expected species slice")
	}
}

func TestPopulation_AddToSpecies(t *testing.T) {
	testSpecies := InitSpecies(5, 3, 1)
	testPop := &Population{}
	testPop.AddToSpecies(testSpecies)
	if len(testPop.GetSpecies()) != 1 {
		t.Fatalf("Expected 1 species in test pop, got %v", len(testPop.GetSpecies()))
	}
	if testPop.GetSpecies()[0] != testSpecies {
		t.Fatalf("Get species did not return expected species slice")
	}
}

func TestPopulation_GetGeneration(t *testing.T) {
	testPop := &Population{generation: 5}
	if testPop.GetGeneration() != testPop.generation {
		t.Fatalf("Expected generation to be %v, but got %v", testPop.generation, testPop.GetGeneration())
	}
}

func TestPopulation_IncrementGeneration(t *testing.T) {
	testPop := &Population{generation: 5}
	testPop.IncrementGeneration()
	if testPop.GetGeneration() != 6 {
		t.Fatalf("Expected generation to be 6, but got %v", testPop.GetGeneration())
	}
}

func TestPopulation_GetTotalInputs(t *testing.T) {
	testPop := &Population{totalInputs: 5}
	if testPop.GetTotalInputs() != testPop.totalInputs {
		t.Fatalf("Expected total inputs to be 5, but got %v", testPop.totalInputs)
	}
}

func TestPopulation_GetTotalOutputs(t *testing.T) {
	testPop := &Population{totalOutputs: 5}
	if testPop.GetTotalOutputs() != testPop.totalOutputs {
		t.Fatalf("Expected total inputs to be 5, but got %v", testPop.totalOutputs)
	}
}

func TestPopulation_GetGrandChampion(t *testing.T) {
	testGenome := &Genome{}
	testPop := &Population{grandChampion: testGenome}
	if testPop.GetGrandChampion() != testPop.grandChampion {
		t.Fatalf("Did not get expected grand champion")
	}
}

func TestPopulation_SetGrandChampion(t *testing.T) {
	testGenome := &Genome{fitness: 5}
	testSpecies := &Species{genomes: []*Genome{testGenome}}
	testPop := &Population{species: []*Species{testSpecies}}
	testPop.SetGrandChampion()
	if testPop.GetGrandChampion() != testGenome {
		t.Fatalf("Did not get expected grand champion")
	}
}

func TestPopulation_GetFitnessCap(t *testing.T) {
	testPop := Population{fitnessCap: 5}
	if testPop.GetFitnessCap() != testPop.fitnessCap {
		t.Fatalf("Expected fitness cap of %v, got %v", testPop.fitnessCap, testPop.GetFitnessCap())
	}
}

func TestPopulation_SetFitnessCap(t *testing.T) {
	testPop := Population{}
	testPop.SetFitnessCap(5)
	if testPop.GetFitnessCap() != 5 {
		t.Fatalf("Expected fitness cap of 5, got %v", testPop.GetFitnessCap())
	}
}

func TestPopulation_GetAllGenomes(t *testing.T) {
	testGenomeOne := &Genome{}
	testGenomeTwo := &Genome{}
	testSpeciesOne := &Species{genomes: []*Genome{testGenomeOne}}
	testSpeciesTwo := &Species{genomes: []*Genome{testGenomeTwo}}
	testPop := Population{species: []*Species{testSpeciesOne, testSpeciesTwo}}
	if len(testPop.GetAllGenomes()) != 2 {
		t.Fatalf("Expected 2 genomes, got %v", len(testPop.GetAllGenomes()))
	}
}
