package main

import (
	"NEAT/Network"
	"log"
)

func main() {
	pop := Network.InitPopulation(5, 3)
	inputs := []float64{-1, -0.5, 0, 0.5, 1}
	for x := 0; x < 100; x++ {
		for i := range pop.GetAllGenomes() {
			if err := pop.GetAllGenomes()[i].TakeInput(inputs); err != nil {
				log.Fatalf("Error: %v", err.Error())
			}
			pop.GetAllGenomes()[i].FeedForward()
			outputSum := 0.0
			for j := range pop.GetAllGenomes()[i].GetOutputs() {
				outputSum = outputSum + pop.GetAllGenomes()[i].GetOutputs()[j]
			}
			pop.GetAllGenomes()[i].SetFitness(outputSum)
		}
		for i := range pop.GetSpecies() {
			pop.GetSpecies()[i].SetChampion()
			/*if pop.GetGeneration() % 5 == 0 {
				pop.GetSpecies()[i].CullTheWeak()
			}*/
		}
		pop.SetGrandChampion()
		log.Println(pop.GetGrandChampion().GetFitness(), pop.GetGeneration())
		pop.Mutate()
	}
	Network.DrawGenome(pop.GetGrandChampion())
}
