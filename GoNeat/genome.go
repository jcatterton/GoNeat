package GoNeat

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Genome struct {
	nodes             []*Node
	connections       []*Connection
	layers            int
	innovationCounter int
	fitness           float64
	mutable           bool
}

func CreateGenome(nodes []*Node, connections []*Connection, layers int, innovation int, fitness float64, mutable bool) *Genome {
	return &Genome{nodes, connections, layers, innovation, fitness, mutable}
}

func InitGenome(input int, output int) *Genome {
	initNodes := []*Node{}
	innovationCounter := 0
	for i := 0; i < input; i++ {
		initNodes = append(initNodes, &Node{layer: 1, weight: 1, innovationNumber: innovationCounter})
		innovationCounter = innovationCounter + 1
	}
	for i := 0; i < output; i++ {
		initNodes = append(initNodes, &Node{layer: 2, innovationNumber: innovationCounter})
		innovationCounter = innovationCounter + 1
	}
	return &Genome{nodes: initNodes, connections: nil, layers: 2, innovationCounter: innovationCounter,
		fitness: 0, mutable: true}
}

func (g *Genome) GetNodes() []*Node {
	return g.nodes
}

func (g *Genome) GetHiddenNodes() []*Node {
	hiddenNodes := []*Node{}
	for i := range g.GetNodes() {
		if g.GetNodes()[i].GetLayer() != 1 && g.GetNodes()[i].GetLayer() != g.GetLayers() {
			hiddenNodes = append(hiddenNodes, g.GetNodes()[i])
		}
	}
	return hiddenNodes
}

func (g *Genome) GetNodesWithLayerLessThan(l int) []*Node {
	layerNodes := []*Node{}
	for i := range g.GetNodes() {
		if g.GetNodes()[i].GetLayer() < l {
			layerNodes = append(layerNodes, g.GetNodes()[i])
		}
	}
	return layerNodes
}

func (g *Genome) GetNodesWithLayerGreaterThan(l int) []*Node {
	layerNodes := []*Node{}
	for i := range g.GetNodes() {
		if g.GetNodes()[i].GetLayer() > l {
			layerNodes = append(layerNodes, g.GetNodes()[i])
		}
	}
	return layerNodes
}

func (g *Genome) AddNode(n *Node) {
	g.IncrementInnovation()
	n.SetInnovationNumber(g.GetInnovation())
	g.nodes = append(g.nodes, n)
}

func (g *Genome) GetConnections() []*Connection {
	return g.connections
}

func (g *Genome) AddConnection(c *Connection) {
	g.IncrementInnovation()
	c.SetInnovationNumber(g.GetInnovation())
	g.connections = append(g.connections, c)
}

func (g *Genome) AddNodeWithoutIncrement(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Genome) AddConnectionWithoutIncrement(c *Connection) {
	g.connections = append(g.connections, c)
}

func (g *Genome) GetLayers() int {
	return g.layers
}

func (g *Genome) IncrementLayers() {
	g.layers = g.layers + 1
}

func (g *Genome) DecrementLayers() {
	g.layers = g.layers - 1
}

func (g *Genome) SetLayers(l int) {
	g.layers = l
}

func (g *Genome) AddRandomNode() {
	if len(g.GetConnections()) == 0 {
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())
	randomConnection := g.GetConnections()[rand.Intn(len(g.GetConnections()))]
	newNode := Node{innovationNumber: g.GetInnovation() + 1}
	newConnection := Connection{innovationNumber: g.GetInnovation() + 2, weight: (rand.Float64() * 2) - 1}

	if randomConnection.GetNodeB().GetLayer()-randomConnection.GetNodeA().GetLayer() == 1 {
		tempLayer := randomConnection.GetNodeB().GetLayer()
		for i := range g.GetNodes() {
			if g.GetNodes()[i].GetLayer() >= tempLayer {
				g.GetNodes()[i].SetLayer(g.GetNodes()[i].GetLayer() + 1)
			}
		}
		g.IncrementLayers()
	} else if randomConnection.GetNodeB().GetLayer()-randomConnection.GetNodeA().GetLayer() < 1 {
		log.Fatalf("Connection has NodeA with layer %v which is not less than NodeB with layer %v",
			randomConnection.GetNodeA().GetLayer(), randomConnection.GetNodeB().GetLayer())
	}
	newNode.SetLayer((randomConnection.GetNodeB().GetLayer() + randomConnection.GetNodeA().GetLayer()) / 2)

	tempNode := &*randomConnection.GetNodeB()
	randomConnection.SetNodeB(&newNode)
	tempNode.RemoveFromInwardConnections(randomConnection)
	newNode.AddToInwardConnections(randomConnection)
	newConnection.SetNodeA(&newNode)
	newNode.AddToOutwardConnections(&newConnection)
	newConnection.SetNodeB(tempNode)
	tempNode.AddToInwardConnections(&newConnection)

	g.AddNode(&newNode)
	g.AddConnection(&newConnection)

	newNode.Deactivate()
}

func (g *Genome) AddRandomConnection() {
	if g.IsFullyConnected() {
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	randomNodeA := g.GetNodes()[rand.Intn(len(g.GetNodes()))]
	randomNodeB := g.GetNodes()[rand.Intn(len(g.GetNodes()))]
	for (randomNodeA.GetLayer() >= randomNodeB.GetLayer()) || randomNodeA.IsConnectedTo(randomNodeB) {
		rand.Seed(time.Now().UTC().UnixNano())
		randomNodeA = g.GetNodes()[rand.Intn(len(g.GetNodes()))]
		rand.Seed(time.Now().UTC().UnixNano())
		randomNodeB = g.GetNodes()[rand.Intn(len(g.GetNodes()))]
	}

	newConnection := Connection{nodeA: randomNodeA, nodeB: randomNodeB, weight: rand.Float64(),
		innovationNumber: g.GetInnovation() + 1}
	randomNodeA.AddToOutwardConnections(&newConnection)
	randomNodeB.AddToInwardConnections(&newConnection)
	g.AddConnection(&newConnection)
}

func (g *Genome) IsFullyConnected() bool {
	for i := range g.GetNodes() {
		for j := range g.GetNodes() {
			if g.GetNodes()[i].GetLayer() != g.GetNodes()[j].GetLayer() && !g.GetNodes()[i].IsConnectedTo(g.GetNodes()[j]) {
				return false
			}
		}
	}
	return true
}

func (g *Genome) GetNodesWithLayer(l int) []*Node {
	if l > g.GetLayers() {
		return nil
	}

	nodesOnLayer := []*Node{}
	for i := range g.GetNodes() {
		if g.GetNodes()[i].GetLayer() == l {
			nodesOnLayer = append(nodesOnLayer, g.GetNodes()[i])
		}
	}
	return nodesOnLayer
}

func (g *Genome) SortNodesByLayer() {
	sort.Slice(g.nodes, func(i, j int) bool {
		return g.GetNodes()[i].GetLayer() < g.GetNodes()[j].GetLayer()
	})
}

func (g *Genome) FeedForward() {
	g.SortNodesByLayer()
	linearCombination := 0.0

	for i := range g.GetNodes() {
		if g.GetNodes()[i].GetLayer() != 1 {
			for j := range g.GetNodes()[i].GetInwardConnections() {
				linearCombination = linearCombination + (g.GetNodes()[i].GetInwardConnections()[j].GetWeight() *
					g.GetNodes()[i].GetInwardConnections()[j].GetNodeA().GetWeight())
			}
			g.GetNodes()[i].SetWeight(linearCombination)
			g.GetNodes()[i].Sigmoid()
		}
	}
	linearCombination = 0
}

func (g *Genome) GetInnovation() int {
	return g.innovationCounter
}

func (g *Genome) IncrementInnovation() {
	g.innovationCounter = g.innovationCounter + 1
}

func (g *Genome) SetInnovationCounter(i int) {
	g.innovationCounter = i
}

func (g *Genome) GetFitness() float64 {
	return g.fitness
}

func (g *Genome) SetFitness(f float64) {
	g.fitness = f
}

func (g *Genome) Clone() *Genome {
	newNodes := []*Node{}
	newConnections := []*Connection{}

	if len(g.GetNodes()) > 0 {
		for i := range g.GetNodes() {
			newNodes = append(newNodes, g.GetNodes()[i].Clone())
		}
	}
	if len(g.GetConnections()) > 0 {
		for i := range g.GetConnections() {
			newConnections = append(newConnections, g.GetConnections()[i].Clone())
		}
		for i := range g.GetConnections() {
			newConnections[i].SetNodeA(newNodes[NodeIndex(g.GetNodes(), g.GetConnections()[i].GetNodeA())])
			newConnections[i].SetNodeB(newNodes[NodeIndex(g.GetNodes(), g.GetConnections()[i].GetNodeB())])
		}
	}
	if len(g.GetNodes()) > 0 {
		for i := range g.GetNodes() {
			if len(g.GetNodes()[i].GetOutwardConnections()) > 0 {
				for j := range g.GetNodes()[i].GetOutwardConnections() {
					newNodes[i].AddToOutwardConnections(newConnections[ConnectionIndex(g.GetConnections(),
						g.GetNodes()[i].GetOutwardConnections()[j])])
				}
			}
			if len(g.GetNodes()[i].GetInwardConnections()) > 0 {
				for j := range g.GetNodes()[i].GetInwardConnections() {
					newNodes[i].AddToInwardConnections(newConnections[ConnectionIndex(g.GetConnections(),
						g.GetNodes()[i].GetInwardConnections()[j])])
				}
			}
		}
	}

	newGenome := &Genome{nodes: newNodes, connections: newConnections, layers: g.GetLayers(),
		innovationCounter: g.GetInnovation(), mutable: g.IsMutable()}

	return newGenome
}

func NodeIndex(nodes []*Node, node *Node) int {
	for i := range nodes {
		if nodes[i] == node {
			return i
		}
	}
	return -1
}

func NodeInnovationIndex(nodes []*Node, node *Node) int {
	for i := range nodes {
		if nodes[i].GetInnovationNumber() == node.GetInnovationNumber() {
			return i
		}
	}
	return -1
}

func ConnectionIndex(connections []*Connection, connection *Connection) int {
	for i := range connections {
		if connections[i] == connection {
			return i
		}
	}
	return -1
}

func ConnectionInnovationIndex(connections []*Connection, connection *Connection) int {
	for i := range connections {
		if connections[i].GetInnovationNumber() == connection.GetInnovationNumber() {
			return i
		}
	}
	return -1
}

func (g *Genome) TakeInput(input []float64) error {
	inputNodes := g.GetNodesWithLayer(1)
	sort.Slice(inputNodes, func(i, j int) bool {
		return inputNodes[i].GetInnovationNumber() < inputNodes[j].GetInnovationNumber()
	})

	if len(input) != len(inputNodes) {
		return errors.New(fmt.Sprintf("Invalid number of inputs. Got %v, expected %v.",
			len(input), len(inputNodes)))
	}
	for i := range inputNodes {
		g.GetNodes()[NodeInnovationIndex(g.GetNodes(), inputNodes[i])].SetWeight(input[i])
		g.GetNodes()[NodeInnovationIndex(g.GetNodes(), inputNodes[i])].Activate()
	}
	return nil
}

func (g *Genome) Mutate() int {
	if !g.IsMutable() {
		return 0
	}

	r := rand.Float64()
	if r <= 0.8 {
		for i := range g.GetConnections() {
			rand.Seed(time.Now().UTC().UnixNano())
			if rand.Float64() >= 0.1 {
				if rand.Float64() >= 0.9 {
					rand.Seed(time.Now().UTC().UnixNano())
					g.GetConnections()[i].SetWeight((rand.Float64() * 2) - 1)
				} else {
					g.GetConnections()[i].SetWeight((rand.Float64() * (0.2 * g.GetConnections()[i].GetWeight())) - 0.1*g.GetConnections()[i].GetWeight())
				}
			}
		}
		return 0
	} else if r > 0.8 && r <= 0.9 {
		rand.Seed(time.Now().UTC().UnixNano())
		if rand.Float64() >= 0.1 {
			g.AddRandomConnection()
			return 1
		}
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
		if rand.Float64() >= 0.1 {
			g.AddRandomNode()
			return 2
		}
	}
	return 0
}

func (g *Genome) GetOutputs() []float64 {
	outputs := []float64{}
	outputNodes := g.GetNodesWithLayer(g.GetLayers())

	for i := range outputNodes {
		outputs = append(outputs, outputNodes[i].GetWeight())
	}

	return outputs
}

func (g *Genome) RemoveDisjointConnections() {
	for i := range g.GetConnections() {
		if g.GetConnections()[i].GetNodeA() == nil || g.GetConnections()[i].GetNodeB() == nil ||
			NodeInnovationIndex(g.GetNodes(), g.GetConnections()[i].GetNodeA()) == -1 ||
			NodeInnovationIndex(g.GetNodes(), g.GetConnections()[i].GetNodeB()) == -1 {
			copy(g.connections[i:], g.connections[i+1:])
			g.connections = g.connections[:len(g.connections)-1]
		}
	}
}

func (g *Genome) HandleDisjointNodes() {
	for i := range g.GetNodes() {
		newConnection := &Connection{}
		if g.GetNodes()[i].GetLayer() != 1 && len(g.GetNodes()[i].GetInwardConnections()) == 0 {
			newConnection = &Connection{
				nodeA: g.GetNodesWithLayerLessThan(g.GetNodes()[i].
					GetLayer())[rand.Intn(len(g.GetNodesWithLayerLessThan(g.GetNodes()[i].GetLayer())))],
				nodeB:  g.GetNodes()[i],
				weight: (rand.Float64() * 2) - 1,
			}
			newConnection.GetNodeA().AddToOutwardConnections(newConnection)
			newConnection.GetNodeB().AddToInwardConnections(newConnection)
			g.AddConnection(newConnection)
		}
		if g.GetNodes()[i].GetLayer() != g.GetLayers() && len(g.GetNodes()[i].GetOutwardConnections()) == 0 {
			newConnection = &Connection{
				nodeA: g.GetNodes()[i],
				nodeB: g.GetNodesWithLayerGreaterThan(g.GetNodes()[i].
					GetLayer())[rand.Intn(len(g.GetNodesWithLayerGreaterThan(g.GetNodes()[i].GetLayer())))],
				weight: (rand.Float64() * 2) - 1,
			}
			newConnection.GetNodeA().AddToOutwardConnections(newConnection)
			newConnection.GetNodeB().AddToInwardConnections(newConnection)
			g.AddConnection(newConnection)
		}
	}
}

func (g *Genome) IsMutable() bool {
	return g.mutable
}

func (g *Genome) SetMutability(m bool) {
	g.mutable = m
}

func (g *Genome) CompatibleWith(gen *Genome) bool {
	excessAndDisjoint, averageWeightDiff := g.getExcessDisjointAvgWeightDiff(gen)

	largeGenomeNormaliser := len(g.nodes) + len(g.connections) - 20
	if largeGenomeNormaliser < 1 {
		largeGenomeNormaliser = 1
	}

	excessCoef := 1.0
	weightDiffCoef := 0.5
	compatibilityThreshold := 3.0

	compatibility := (excessCoef * excessAndDisjoint / float64(largeGenomeNormaliser)) + (weightDiffCoef * averageWeightDiff)
	return compatibilityThreshold > compatibility
}

func (g *Genome) getExcessDisjointAvgWeightDiff(gen *Genome) (float64, float64) {
	if len(g.nodes) == 0 || len(g.connections) == 0 || len(gen.nodes) == 0 || len(gen.connections) == 0 {
		return 0, 0
	}

	matching := 0
	totalDiff := 0.0
	for i := range g.nodes {
		for j := range gen.nodes {
			if g.nodes[i].innovationNumber == gen.nodes[j].innovationNumber {
				matching++
				totalDiff = totalDiff + math.Abs(g.nodes[i].weight-gen.nodes[j].weight)
				break
			}
		}
	}
	for i := range g.connections {
		for j := range gen.connections {
			if g.connections[i].innovationNumber == gen.connections[j].innovationNumber {
				matching++
				totalDiff = totalDiff + math.Abs(g.connections[i].weight-gen.connections[j].weight)
				break
			}
		}
	}
	if matching == 0 {
		return float64(len(g.nodes) + len(g.connections) + len(gen.nodes) + len(gen.connections) - 2*matching), 100
	}

	return float64(len(g.nodes) + len(g.connections) + len(gen.nodes) + len(gen.connections) - 2*matching), totalDiff / float64(matching)
}
