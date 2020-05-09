package GoNeat

import (
	"math"
	"testing"
)

func TestNode_IsActivated(t *testing.T) {
	testNode := Node{activated: false}
	if testNode.IsActivated() != testNode.activated {
		t.Fatal("Expected isActivated to be false, got true.")
	}
}

func TestNode_Activate(t *testing.T) {
	testNode := Node{}
	testNode.Activate()
	if testNode.IsActivated() != true {
		t.Fatal("Expected isActivated to be true, got false.")
	}
}

func TestNode_Deactivate(t *testing.T) {
	testNode := Node{}
	testNode.Deactivate()
	if testNode.IsActivated() != false {
		t.Fatal("Expected isActivated to be false, go true.")
	}
}

func TestNode_GetOutwardConnections(t *testing.T) {
	testNode := Node{outwardConnections: []*Connection{&Connection{}}}
	if len(testNode.GetOutwardConnections()) != 1 {
		t.Fatalf("Expected slice of length 1, got slice of length %v", len(testNode.GetOutwardConnections()))
	}
}

func TestNode_GetInwardConnections(t *testing.T) {
	testNode := Node{inwardConnections: []*Connection{&Connection{}}}
	if len(testNode.GetInwardConnections()) != 1 {
		t.Fatalf("Expected slice of length 1, got slice of length %v.", len(testNode.GetInwardConnections()))
	}
}

func TestNode_AddToOutwardConnections(t *testing.T) {
	testNode := Node{}
	testConnection := Connection{}
	testNode.AddToOutwardConnections(&testConnection)
	if len(testNode.GetOutwardConnections()) != 1 {
		t.Fatalf("Expected slice of length 1, got slice of length %v.", len(testNode.GetOutwardConnections()))
	}
	if &testConnection != testNode.GetOutwardConnections()[0] {
		t.Fatalf("Outward connection in test node does not match test conneciton.")
	}
}

func TestNode_AddToInwardConnections(t *testing.T) {
	testNode := Node{}
	testConnection := Connection{}
	testNode.AddToInwardConnections(&testConnection)
	if len(testNode.GetInwardConnections()) != 1 {
		t.Fatalf("Expected slice of length 1, got slice of length %v", len(testNode.GetInwardConnections()))
	}
	if &testConnection != testNode.GetInwardConnections()[0] {
		t.Fatalf("Outward connection in test node does not match test conneciton.")
	}
}

func TestNode_RemoveFromOutwardConnections(t *testing.T) {
	testConnection := Connection{}
	testNode := Node{outwardConnections: []*Connection{&testConnection}}
	testNode.RemoveFromOutwardConnections(&testConnection)
	if len(testNode.GetOutwardConnections()) != 0 {
		t.Fatalf("Expected slice of length 0, got slice of length %v", len(testNode.GetOutwardConnections()))
	}

	unaddedConnection := Connection{}
	if testNode.RemoveFromOutwardConnections(&unaddedConnection) != -1 {
		t.Fatal("Expected -1 response when attempting to remove connection which has not been added to genome.")
	}
}

func TestNode_RemoveFromInwardConnections(t *testing.T) {
	testConnection := Connection{}
	testNode := Node{inwardConnections: []*Connection{&testConnection}}
	testNode.RemoveFromInwardConnections(&testConnection)
	if len(testNode.GetInwardConnections()) != 0 {
		t.Fatalf("Expected slice of lenght 0, got slice of length %v", len(testNode.GetInwardConnections()))
	}

	unaddedConnection := Connection{}
	if testNode.RemoveFromInwardConnections(&unaddedConnection) != -1 {
		t.Fatal("Expected -1 response when attempting to remove connection which has not been added to genome.")
	}
}

func TestNode_RemoveFromInwardConnectionsUsingTempNode(t *testing.T) {
	testConnection := Connection{}
	testNode := Node{inwardConnections: []*Connection{&testConnection}}
	tempNode := &testNode
	tempNode.RemoveFromInwardConnections(&testConnection)
	if len(testNode.GetInwardConnections()) != 0 {
		t.Fatal("Expected test connection be have been removed, but it was not.")
	}
}

func TestNode_ClearOutwardConnections(t *testing.T) {
	testConnection := Connection{}
	testNode := Node{outwardConnections: []*Connection{&testConnection}}
	testNode.ClearOutwardConnections()
	if len(testNode.GetOutwardConnections()) != 0 {
		t.Fatalf("Expected outward connections to be cleared, but still has length %v",
			len(testNode.GetOutwardConnections()))
	}
}

func TestNode_ClearInwardConnections(t *testing.T) {
	testConnection := Connection{}
	testNode := Node{inwardConnections: []*Connection{&testConnection}}
	testNode.ClearInwardConnections()
	if len(testNode.GetInwardConnections()) != 0 {
		t.Fatalf("Expected inward connections to be cleared, but still has length %v",
			len(testNode.GetInwardConnections()))
	}
}

func TestNode_GetLayer(t *testing.T) {
	testNode := Node{layer: 3}
	if testNode.GetLayer() != testNode.layer {
		t.Fatalf("Expected test node to have layer 3, got layer %v", testNode.layer)
	}
}

func TestNode_SetLayer(t *testing.T) {
	testNode := Node{}
	testNode.SetLayer(3)
	if testNode.GetLayer() != 3 {
		t.Fatalf("Expected test node to have layer 3, got layer %v", testNode.layer)
	}
}

func TestNode_IsConnectedTo(t *testing.T) {
	testNodeA := Node{}
	testNodeB := Node{}
	testConnection := &Connection{nodeA: &testNodeA, nodeB: &testNodeB}
	testNodeA.AddToOutwardConnections(testConnection)
	testNodeB.AddToInwardConnections(testConnection)
	if !testNodeA.IsConnectedTo(&testNodeB) {
		t.Fatal("Expected testNodeA to connect to testNodeB, but it was not.")
	}
	if !testNodeB.IsConnectedTo(&testNodeA) {
		t.Fatal("Expected testNodeB to connect to testNodeA, but it was not.")
	}

	testConnection.SetNodeA(nil)
	testConnection.SetNodeB(nil)
	if testConnection.GetNodeA() != nil {
		t.Fatalf("Did not successfully set connection node A to nil.")
	}
	if testConnection.GetNodeB() != nil {
		t.Fatalf("Did not successfully set connection node A to nil.")
	}

	testNodeA.RemoveFromOutwardConnections(testConnection)
	testNodeB.RemoveFromInwardConnections(testConnection)
	if len(testNodeA.GetOutwardConnections()) != 0 {
		t.Fatalf("Did not successfully remove connection from node A outward connections")
	}
	if len(testNodeB.GetInwardConnections()) != 0 {
		t.Fatalf("Did not successfully remove connection from node B inward connections")
	}

	if testNodeA.IsConnectedTo(&testNodeB) {
		t.Fatalf("Expected testNodeA to not connect to testNodeB, but it did.")
	}
}

func TestNode_GetWeight(t *testing.T) {
	testNodeA := Node{weight: 0.51}
	if testNodeA.GetWeight() != testNodeA.weight {
		t.Fatalf("Expected test node to have weight %v, but got %v", testNodeA.GetWeight(), testNodeA.weight)
	}
}

func TestNode_SetWeight(t *testing.T) {
	testNodeA := Node{}
	testNodeA.SetWeight(0.51)
	if testNodeA.GetWeight() != 0.51 {
		t.Fatalf("Expected test node to have weight 0.51, but got %v", testNodeA.GetWeight())
	}
}

func TestNode_Sigmoid(t *testing.T) {
	testNodeA := Node{weight: -1}
	testNodeB := Node{weight: 1}

	testNodeA.Sigmoid()
	testNodeB.Sigmoid()

	if testNodeA.IsActivated() {
		t.Fatalf("Expected test node to be deactivated, but it is activated. Weight is %v, sigmoid value is %v",
			testNodeA.GetWeight(), 1/(1+math.Exp(testNodeA.GetWeight()*-4.9)))
	}
	if !testNodeB.IsActivated() {
		t.Fatalf("Expected test node to be activated, but it is deactivated. Weight is %v, sigmoid value is %v",
			testNodeB.GetWeight(), 1/(1+math.Exp(testNodeB.GetWeight()*-4.9)))
	}
}

func TestNode_GetInnovationNumber(t *testing.T) {
	testNode := Node{innovationNumber: 5}
	if testNode.GetInnovationNumber() != testNode.innovationNumber {
		t.Fatalf("Expected test node to have innovation number 5, got %v", testNode.GetInnovationNumber())
	}
}

func TestNode_SetInnovationNumber(t *testing.T) {
	testNode := Node{}
	testNode.SetInnovationNumber(5)
	if testNode.GetInnovationNumber() != 5 {
		t.Fatalf("Expected test node to have innovation number 5, got %v", testNode.GetInnovationNumber())
	}
}

func TestNode_Clone(t *testing.T) {
	testNode := Node{activated: true, layer: 3, outwardConnections: []*Connection{{}},
		inwardConnections: []*Connection{{}}, weight: 0.75, innovationNumber: 6}
	copyNode := testNode.Clone()

	if copyNode.IsActivated() != testNode.IsActivated() {
		t.Fatalf("Expected copied node to have the same activation status as test node, but copied node has "+
			"activation status %v, and test node has activation status %v.",
			copyNode.IsActivated(), testNode.IsActivated())
	}
	if copyNode.GetLayer() != testNode.GetLayer() {
		t.Fatalf("Expected copied node to have the same layer as test node, but copied node has layer %v and "+
			"test node has layer %v.", copyNode.GetLayer(), testNode.GetLayer())
	}
	if copyNode.GetWeight() != testNode.GetWeight() {
		t.Fatalf("Expected copied node to have the same weight as test node, but copied node has weight %v "+
			"and test node has weight %v.", copyNode.GetWeight(), testNode.GetWeight())
	}
	if copyNode.GetInnovationNumber() != testNode.GetInnovationNumber() {
		t.Fatalf("Expected copied node to have the same innovation number as test node, but copied node has "+
			"innovation number %v and test node has innovation number %v", copyNode.GetInnovationNumber(),
			testNode.GetInnovationNumber())
	}
	if copyNode == &testNode {
		t.Fatalf("Copied node and test node have the same memory addresses, %v and %v", copyNode, &testNode)
	}
}
