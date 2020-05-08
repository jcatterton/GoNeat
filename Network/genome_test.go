package Network

import (
	"testing"
)

func TestInitGenome(t *testing.T) {
	testGenome := InitGenome(5, 3)

	layerOneNodes := 0
	layerTwoNodes := 0
	layerThreeNodes := 0
	for i := range testGenome.GetNodes() {
		if testGenome.GetNodes()[i].GetLayer() == 1 {
			layerOneNodes = layerOneNodes + 1
		} else if testGenome.GetNodes()[i].GetLayer() == 2 {
			layerTwoNodes = layerTwoNodes + 1
		} else if testGenome.GetNodes()[i].GetLayer() == 3 {
			layerThreeNodes = layerThreeNodes + 1
		}
	}

	if testGenome.GetLayers() != 3 {
		t.Fatalf("Expected 3 layers, got %v", testGenome.GetLayers())
	}
	if layerOneNodes != 5 {
		t.Fatalf("Expected layer one to have five nodes, got %v", layerOneNodes)
	}
	if layerTwoNodes != (layerOneNodes+layerThreeNodes)/2 {
		t.Fatalf("Expected layer two to have %v nodes, got %v", (layerOneNodes+layerThreeNodes)/2,
			layerTwoNodes)
	}
	if layerThreeNodes != 3 {
		t.Fatalf("Expected layer three to have three nodes, got %v", layerThreeNodes)
	}
	if len(testGenome.GetConnections()) != ((layerOneNodes * layerTwoNodes) + (layerTwoNodes * layerThreeNodes)) {
		t.Fatalf("Expected %v connections, got %v",
			(layerOneNodes*layerTwoNodes)+(layerTwoNodes*layerThreeNodes),
			len(testGenome.GetConnections()))
	}
	for i := 0; i < len(testGenome.GetNodes())-1; i++ {
		if testGenome.GetNodes()[i].GetInnovationNumber() == testGenome.GetNodes()[i+1].GetInnovationNumber() {
			t.Fatalf("Expected nodes to be initialized with different innovation numbers, but node at index "+
				"%v and %v have the same innovation number of %v", i, i+1,
				testGenome.GetNodes()[i].GetInnovationNumber())
		}
	}
	for i := 0; i < len(testGenome.GetConnections())-1; i++ {
		if testGenome.GetConnections()[i].GetInnovationNumber() ==
			testGenome.GetConnections()[i+1].GetInnovationNumber() {
			t.Fatalf("Expected connections to be initialized with different innovation numbers, but "+
				"connections at index %v and %v have the same innovation number of %v", i, i+1,
				testGenome.GetConnections()[i].GetInnovationNumber())
		}
	}
}

func TestGenome_GetNodes(t *testing.T) {
	testGenome := Genome{nodes: []*Node{&Node{}}}
	if len(testGenome.GetNodes()) != 1 {
		t.Fatalf("Expected test genome to have node slice of length 1, got slice of length %v",
			len(testGenome.GetNodes()))
	}
}

func TestGenome_GetHiddenNodes(t *testing.T) {
	testNodeOne := &Node{layer: 1}
	testNodeTwo := &Node{layer: 2}
	testNodeThree := &Node{layer: 3}
	testGenome := Genome{nodes: []*Node{testNodeOne, testNodeTwo, testNodeThree}, layers: 3}
	if len(testGenome.GetHiddenNodes()) != 1 {
		t.Fatalf("Expected one hidden node, got %v", len(testGenome.GetHiddenNodes()))
	}
	if testGenome.GetHiddenNodes()[0] != testNodeTwo {
		t.Fatalf("Expected test node two to be in hidden nodes slice, but it was not. Got %v",
			testGenome.GetHiddenNodes()[0])
	}
}

func TestGenome_GetNodesWithLayerGreaterThan(t *testing.T) {
	testNodeOne := &Node{layer: 1}
	testNodeTwo := &Node{layer: 2}
	testGenome := Genome{nodes: []*Node{testNodeOne, testNodeTwo}}
	if len(testGenome.GetNodesWithLayerGreaterThan(1)) != 1 {
		t.Fatalf("Expected one node, got %v", len(testGenome.GetNodesWithLayerGreaterThan(1)))
	}
	if testGenome.GetNodesWithLayerGreaterThan(1)[0] != testNodeTwo {
		t.Fatalf("Expected test node two to be first element in slice, but it was not")
	}
}

func TestGenome_GetNodesWithLayerLessThan(t *testing.T) {
	testNodeOne := &Node{layer: 1}
	testNodeTwo := &Node{layer: 2}
	testGenome := Genome{nodes: []*Node{testNodeOne, testNodeTwo}}
	if len(testGenome.GetNodesWithLayerLessThan(2)) != 1 {
		t.Fatalf("Expected one node, got %v", len(testGenome.GetNodesWithLayerLessThan(2)))
	}
	if testGenome.GetNodesWithLayerLessThan(2)[0] != testNodeOne {
		t.Fatalf("Expected test node one to be first element in slice, but it was not")
	}
}

func TestGenome_AddNode(t *testing.T) {
	testGenome := Genome{}
	testNode := Node{}
	testGenome.AddNode(&testNode)
	if len(testGenome.GetNodes()) != 1 {
		t.Fatalf("Expected test genome to have node slice of length 1, got slice of length %v",
			len(testGenome.GetNodes()))
	}
	if testGenome.GetInnovation() != 1 {
		t.Fatalf("Expected test genome innovation to have incremented to 1, but it was not.")
	}
	if testNode.GetInnovationNumber() != testGenome.GetInnovation() {
		t.Fatalf("Expected added node and genome to have the same innovation number, but node has innovation "+
			"number %v and genome has innovation number %v", testNode.GetInnovationNumber(), testGenome.GetInnovation())
	}
}

func TestGenome_GetConnections(t *testing.T) {
	testGenome := Genome{connections: []*Connection{&Connection{}}}
	if len(testGenome.GetConnections()) != 1 {
		t.Fatalf("Expected test genome to have connection slice of length 1, got slice of length %v",
			len(testGenome.GetConnections()))
	}
}

func TestGenome_AddConnection(t *testing.T) {
	testGenome := Genome{}
	testConnection := Connection{}
	testGenome.AddConnection(&testConnection)
	if len(testGenome.GetConnections()) != 1 {
		t.Fatalf("Expected test genome to have connection slice of length 1, got slice of length %v",
			len(testGenome.GetConnections()))
	}
	if testGenome.GetInnovation() != 1 {
		t.Fatalf("Expected test genome innovation to have incremented to 1, but it was not.")
	}
	if testConnection.GetInnovationNumber() != testGenome.GetInnovation() {
		t.Fatalf("Expected added connection and genome to have the same innovation number, but connection has "+
			"innovation %v and genome has innovation %v",
			testConnection.GetInnovationNumber(), testGenome.GetInnovation())
	}
}

func TestGenome_GetLayers(t *testing.T) {
	testGenome := Genome{layers: 3}
	if testGenome.GetLayers() != testGenome.layers {
		t.Fatalf("Expected test genome to have 3 layers, got %v layers", testGenome.layers)
	}
}

func TestGenome_IncrementLayers(t *testing.T) {
	testGenome := Genome{layers: 3}
	testGenome.IncrementLayers()
	if testGenome.GetLayers() != 4 {
		t.Fatalf("Expected test genome to have 4 layers, got %v layers", testGenome.layers)
	}
}

func TestGenome_DecrementLayers(t *testing.T) {
	testGenome := Genome{layers: 3}
	testGenome.DecrementLayers()
	if testGenome.GetLayers() != 2 {
		t.Fatalf("Expected test genome to have 2 layers, got %v layers", testGenome.layers)
	}
}

func TestGenome_SetLayers(t *testing.T) {
	testGenome := Genome{layers: 0}
	testGenome.SetLayers(3)
	if testGenome.GetLayers() != 3 {
		t.Fatalf("Expected testg enome to have 3 layers, got %v layers", testGenome.layers)
	}
}

func TestGenome_AddRandomNode(t *testing.T) {
	testGenome := InitGenome(5, 3)

	previousLength := len(testGenome.GetNodes())
	testGenome.AddRandomNode()
	newestNode := testGenome.GetNodes()[len(testGenome.GetNodes())-1]

	if len(testGenome.GetNodes()) != previousLength+1 {
		t.Fatalf("Failed adding random node: previous number of nodes was %v, current number is %v",
			previousLength, len(testGenome.GetNodes()))
	}
	if newestNode.GetLayer() >= testGenome.GetLayers() && newestNode.GetLayer() != 1 {
		t.Fatalf("Node added at invalid layer %v, final layer is layer %v", newestNode.GetLayer(),
			testGenome.GetLayers())
	}
	if len(newestNode.GetInwardConnections()) != 1 && len(newestNode.GetOutwardConnections()) != 1 {
		t.Fatalf("Expected one inward and one outward connection for new node, instead got %v inward"+
			"and %v outward connections", newestNode.GetInwardConnections(), newestNode.GetOutwardConnections())
	}
	if newestNode.GetInwardConnections()[0].GetNodeA().GetLayer() >= newestNode.GetLayer() ||
		newestNode.GetOutwardConnections()[0].GetNodeB().GetLayer() <= newestNode.GetLayer() {
		t.Fatalf("Expected new node to connect to node in previous layer and later layer, instead got "+
			"connection to nodes in layers %v and %v, new node is in layer %v",
			newestNode.GetInwardConnections()[0].GetNodeA().GetLayer(),
			newestNode.GetOutwardConnections()[0].GetNodeB().GetLayer(), newestNode.GetLayer())
	}
}

func TestGenome_AddRandomConnection(t *testing.T) {
	testNodes := []*Node{
		{layer: 1},
		{layer: 1},
		{layer: 1},
		{layer: 1},
		{layer: 2},
		{layer: 2},
		{layer: 2},
		{layer: 3},
		{layer: 3},
	}
	testGenome := Genome{nodes: testNodes, layers: 3}

	testGenome.AddRandomConnection()
	newestConnection := testGenome.GetConnections()[0]

	if len(testGenome.GetConnections()) != 1 {
		t.Fatalf("Expected slice connections to have length 1, instead got length %v", testGenome.GetConnections())
	}
	if newestConnection.GetNodeA().GetLayer() >= newestConnection.GetNodeB().GetLayer() {
		t.Fatalf("New connection node A has layer %v, which is greater than node B layer %v",
			newestConnection.GetNodeA().GetLayer(), newestConnection.GetNodeB().GetLayer())
	}
}

func TestGenome_SortNodesByLayer(t *testing.T) {
	testNodes := []*Node{
		{layer: 3},
		{layer: 3},
		{layer: 2},
		{layer: 2},
		{layer: 2},
		{layer: 1},
		{layer: 1},
		{layer: 1},
		{layer: 1},
	}
	testGenome := Genome{nodes: testNodes, layers: 3}

	testGenome.SortNodesByLayer()

	for i := 0; i < len(testGenome.GetNodes())-1; i++ {
		if testGenome.GetNodes()[i].GetLayer() > testGenome.GetNodes()[i+1].GetLayer() {
			t.Fatalf("Expected nodes slice to be ordered by layer, but node at index %v has layer %v,"+
				"while node at index %v has layer %v",
				i, testGenome.GetNodes()[i].GetLayer(), i+1, testGenome.GetNodes()[i+1].GetLayer())
		}
	}
}

func TestGenome_GetNodesWithLayer(t *testing.T) {
	testGenome := InitGenome(5, 3)

	layerOneNodes := testGenome.GetNodesWithLayer(1)
	layerTwoNodes := testGenome.GetNodesWithLayer(2)
	layerThreeNodes := testGenome.GetNodesWithLayer(3)

	if len(layerOneNodes) != 5 {
		t.Fatalf("Expected %v layer one nodes, but got %v", 5, len(layerOneNodes))
	}
	if len(layerTwoNodes) != 4 {
		t.Fatalf("Expected %v layer two nodes, but got %v", 4, len(layerTwoNodes))
	}
	if len(layerThreeNodes) != 3 {
		t.Fatalf("Expected %v layer three nodes, but got %v", 3, len(layerThreeNodes))
	}
	for i := range layerOneNodes {
		if layerOneNodes[i].GetLayer() != 1 {
			t.Fatalf("Node at index %v in layerOneNodes has layer %v", i, layerOneNodes[i].GetLayer())
		}
	}
	for i := range layerTwoNodes {
		if layerTwoNodes[i].GetLayer() != 2 {
			t.Fatalf("Node at index %v in layerTwoNodes has layer %v", i, layerTwoNodes)
		}
	}
	for i := range layerThreeNodes {
		if layerThreeNodes[i].GetLayer() != 3 {
			t.Fatalf("Node at index %v in layerThreeNodes has layer %v", i, layerThreeNodes)
		}
	}
	if testGenome.GetNodesWithLayer(testGenome.GetLayers()+1) != nil {
		t.Fatalf("Expected GetNodesWithLayer to return nil if input is greater than number of layers, but "+
			"returned non-nil when given %v, even though genome has %v layers.",
			testGenome.GetLayers()+1, testGenome.GetLayers())
	}
}

func TestGenome_FeedForward(t *testing.T) {
	testGenome := InitGenome(5, 3)

	finalLayerNodes := testGenome.GetNodesWithLayer(testGenome.GetLayers())

	testGenome.FeedForward()

	for i := range finalLayerNodes {
		if finalLayerNodes[i].GetWeight() == 0 {
			t.Fatalf("Expected final layer nodes to have a non-zero weight between -1 and 1, but node at index "+
				"%v has weight %v", i, finalLayerNodes[i].GetWeight())
		}
	}
}

func TestGenome_GetInnovation(t *testing.T) {
	testGenome := Genome{innovationCounter: 5}
	if testGenome.GetInnovation() != testGenome.innovationCounter {
		t.Fatalf("Expected innovation counter to have value %v, but got %v.", testGenome.innovationCounter,
			testGenome.GetInnovation())
	}
}

func TestGenome_IncrementInnovation(t *testing.T) {
	testGenome := Genome{innovationCounter: 5}
	testGenome.IncrementInnovation()
	if testGenome.GetInnovation() != 6 {
		t.Fatalf("Expected innovation counter to have value %v, but got %v.", testGenome.innovationCounter+1,
			testGenome.GetInnovation())
	}
}

func TestGenome_SetInnovationCounter(t *testing.T) {
	testGenome := Genome{}
	testGenome.SetInnovationCounter(5)
	if testGenome.GetInnovation() != 5 {
		t.Fatalf("Expected test genome to have innovation of 5, but got %v", testGenome.GetInnovation())
	}
}

func TestGenome_Clone(t *testing.T) {
	testGenome := InitGenome(5, 3)
	copyGenome := testGenome.Clone()

	if copyGenome.GetInnovation() != testGenome.GetInnovation() {
		t.Fatalf("Expected copy genome to have the same innovation as test genome, but copy genome has "+
			"innovation value %v and test genome has innovation value %v", testGenome.GetInnovation(),
			copyGenome.GetInnovation())
	}
	if copyGenome.GetLayers() != testGenome.GetLayers() {
		t.Fatalf("Expected copy genome to have the same number of layers as test genome, but copy genome has "+
			"%v while test genome has %v", copyGenome.GetLayers(), testGenome.GetLayers())
	}
	if len(copyGenome.GetNodes()) != len(testGenome.GetNodes()) {
		t.Fatalf("Expected copy genome to have the same number of nodes as test genome, but copy genome has "+
			"%v nodes while test genome has %v", len(copyGenome.GetNodes()), len(testGenome.GetNodes()))
	}
	if len(copyGenome.GetConnections()) != len(testGenome.GetConnections()) {
		t.Fatalf("Expected copy genome to have the same number of connections as test genome, but copy genome "+
			"has %v connectinos while test genome has %v",
			len(copyGenome.GetConnections()), len(testGenome.GetConnections()))
	}
	for i := range copyGenome.GetNodes() {
		if len(copyGenome.GetNodes()[i].GetOutwardConnections()) !=
			len(testGenome.GetNodes()[i].GetOutwardConnections()) {
			t.Fatalf("Expected each node in copy genome to have the same number of outward connections as the "+
				"corresponding node in test genome, but node at index %v has %v outward connections in copy genome, "+
				"while node at index %v has %v outward connections in test genome.",
				i, len(copyGenome.GetNodes()[i].GetOutwardConnections()),
				i, len(testGenome.GetNodes()[i].GetOutwardConnections()))
		}
	}
	for i := range copyGenome.GetNodes() {
		if len(copyGenome.GetNodes()[i].GetInwardConnections()) !=
			len(testGenome.GetNodes()[i].GetInwardConnections()) {
			t.Fatalf("Expected each node in copy genome to have the same number of inward connections as the "+
				"corresponding node in test genome, but node at index %v has %v inward connections in copy genome, "+
				"while node at index %v has %v inward connections in test genome.",
				i, len(copyGenome.GetNodes()[i].GetInwardConnections()),
				i, len(testGenome.GetNodes()[i].GetInwardConnections()))
		}
	}
}

func TestNodeIndex(t *testing.T) {
	testNode := &Node{}
	testNodes := []*Node{testNode}
	if NodeIndex(testNodes, testNode) != 0 {
		t.Fatalf("Expected test node was not in test nodes slice.")
	}

	testNodesTwo := []*Node{{}}
	if NodeIndex(testNodesTwo, testNode) != -1 {
		t.Fatalf("Expected node index to return -1 when searching for node which was not in slice, but got %v",
			NodeIndex(testNodesTwo, testNode))
	}
}

func TestConnectionIndex(t *testing.T) {
	testConnection := &Connection{}
	testConnections := []*Connection{testConnection}
	if ConnectionIndex(testConnections, testConnection) != 0 {
		t.Fatalf("Expected test connection was not in test connections slice.")
	}

	testConnectionsTwo := []*Connection{{}}
	if ConnectionIndex(testConnectionsTwo, testConnection) != -1 {
		t.Fatalf("Expected connection index to return -1 when searching for node which was not in slice, but "+
			"got %v", ConnectionIndex(testConnections, testConnection))
	}
}

func TestGenome_TakeInput(t *testing.T) {
	testGenome := InitGenome(5, 3)
	testInputs := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	_ = testGenome.TakeInput(testInputs)

	inputNodes := testGenome.GetNodesWithLayer(1)
	for i := range inputNodes {
		if inputNodes[i].GetWeight() != testInputs[i] {
			t.Fatalf("Expected input nodes to match given input, but input at index %v has value %v while "+
				"input node at index %v has value %v", i, inputNodes[i].GetWeight(), i, testInputs[i])
		}
	}

	testInputs = []float64{1.0}
	err := testGenome.TakeInput(testInputs)
	if err == nil {
		t.Fatalf("Expected error for giving %v inputs while there are %v input nodes, but got no error.",
			len(testInputs), len(inputNodes))
	}
}

func TestGenome_GetFitness(t *testing.T) {
	testGenome := &Genome{fitness: 0}
	if testGenome.GetFitness() != testGenome.fitness {
		t.Fatalf("Expected get fitness to return %v, but got %v", testGenome.fitness, testGenome.GetFitness())
	}
}

func TestGenome_SetFitness(t *testing.T) {
	testGenome := &Genome{}
	testGenome.SetFitness(1.15)
	if testGenome.GetFitness() != 1.15 {
		t.Fatalf("Expected get fitness to return 1.15, but got %v", testGenome.GetFitness())
	}
}

func TestGenome_GetOutputs(t *testing.T) {
	testGenome := InitGenome(5, 3)
	testGenome.FeedForward()
	if len(testGenome.GetOutputs()) != len(testGenome.GetNodesWithLayer(testGenome.GetLayers())) {
		t.Fatalf("Expected output array to have the same length as final layer size, but output array has "+
			"%v while final layer has %v nodes", len(testGenome.GetOutputs()),
			len(testGenome.GetNodesWithLayer(testGenome.GetLayers())))
	}
}

func TestGenome_Mutate(t *testing.T) {
	testGenome := InitGenome(5, 3)
	testGenomeClone := testGenome.Clone()
	testGenome.SetMutability(true)
	testGenome.Mutate()
	testGenome.FeedForward()
	testGenomeClone.FeedForward()
	testGenomeOutputs := testGenome.GetOutputs()
	testGenomeCloneOutputs := testGenomeClone.GetOutputs()

	genomesDifferent := false
	for i := range testGenomeOutputs {
		genomesDifferent = testGenomeOutputs[i] != testGenomeCloneOutputs[i]
	}
	if !genomesDifferent {
		t.Fatalf("Expected different output arrays after mutating, but both genomes have the same output "+
			"weights: %v and %v", testGenomeOutputs, testGenomeCloneOutputs)
	}
}

func TestGenome_CantMutateIfImmutable(t *testing.T) {
	testGenome := InitGenome(5, 3)
	testGenome.SetMutability(false)
	if testGenome.Mutate() != 0 {
		t.Fatalf("Expected test genome to be unable to mutate, but it did.")
	}
}

func TestGenome_IsFullyConnected(t *testing.T) {
	testGenome := InitGenome(5, 3)
	for i := range testGenome.GetNodes() {
		for j := range testGenome.GetNodes() {
			if testGenome.GetNodes()[j].GetLayer() > testGenome.GetNodes()[i].GetLayer() && !testGenome.GetNodes()[i].IsConnectedTo(testGenome.GetNodes()[j]) {
				testConnection := Connection{}
				testConnection.SetNodeA(testGenome.GetNodes()[i])
				testConnection.SetNodeB(testGenome.GetNodes()[j])
				testGenome.GetNodes()[i].AddToOutwardConnections(&testConnection)
				testGenome.GetNodes()[j].AddToInwardConnections(&testConnection)
				testGenome.AddConnection(&testConnection)
			}
		}
	}

	if !testGenome.IsFullyConnected() {
		t.Fatalf("Expected genome to be fully connected, but it is not.")
	}

	testGenome.AddRandomNode()
	if testGenome.IsFullyConnected() {
		t.Fatalf("Expected genome to not be fully connected, but it is.")
	}
}

func TestGenome_RemoveDisjointConnections(t *testing.T) {
	testConnection := &Connection{}
	testGenome := Genome{connections: []*Connection{testConnection}}
	testGenome.RemoveDisjointConnections()
	if len(testGenome.GetConnections()) != 0 {
		t.Fatalf("Expected connection to be removed, but it was not. Connection slice has length %v",
			len(testGenome.GetConnections()))
	}

	testNodeA := &Node{}
	testNodeB := &Node{}
	testNodeA.AddToOutwardConnections(testConnection)
	testNodeB.AddToInwardConnections(testConnection)
	testConnection.SetNodeA(testNodeA)
	testConnection.SetNodeB(testNodeB)
	testGenome.AddNode(testNodeA)
	testGenome.AddNode(testNodeB)
	testGenome.AddConnection(testConnection)
	testGenome.RemoveDisjointConnections()
	if len(testGenome.GetConnections()) != 1 {
		t.Fatalf("Expected connection to not be removed, but it was.")
	}
}

func TestGenome_HandleDisjointNodes(t *testing.T) {
	testNodeA := &Node{layer: 1}
	testNodeB := &Node{layer: 2}
	testGenome := Genome{nodes: []*Node{testNodeA, testNodeB}, layers: 2}
	testGenome.HandleDisjointNodes()
	if !testNodeA.IsConnectedTo(testNodeB) {
		t.Fatalf("Expected test node a to be connected to test node b, but it was not.")
	}
}

func TestGenome_GetMutability(t *testing.T) {
	testGenome := &Genome{mutable: true}
	if testGenome.IsMutable() != testGenome.mutable {
		t.Fatalf("Expected test genome to be mutable, but it is not.")
	}
}

func TestGenome_SetMutability(t *testing.T) {
	testGenome := &Genome{mutable: true}
	testGenome.SetMutability(false)
	if testGenome.IsMutable() {
		t.Fatalf("Expected test genome to be immutable, but it is not.")
	}
}
