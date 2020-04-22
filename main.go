package main

import (
	"NEAT/Network"
	"log"
)

func main() {
	pop := Network.InitPopulation(3, 5)
	log.Println(len(pop.GetSpecies()))
	log.Println(len(pop.GetSpecies()[0].GetGenomes()))
}
