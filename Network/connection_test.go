package Network

import (
	"testing"
)

func TestConnection_GetWeight(t *testing.T) {
	c := Connection{weight: 0.5}
	if c.weight != 0.5 {
		t.Fatalf("Expected weight of 0.5, got weight of %v", c.weight)
	}
}

func TestConnection_SetWeight(t *testing.T) {
	c := Connection{}
	c.SetWeight(0.5)
	if c.weight != 0.5 {
		t.Fatalf("Expected weight of 0.5, got weight of %v", c.weight)
	}
}

func TestConnection_GetNodeA(t *testing.T) {
	n := Node{}
	c := Connection{nodeA: &n}
	if c.GetNodeA() != &n {
		t.Fatal("Node A did not match expected Node A")
	}
}

func TestConnection_GetNodeB(t *testing.T) {
	n := Node{}
	c := Connection{nodeB: &n}
	if c.GetNodeB() != &n {
		t.Fatal("Node B did not match expected Node B")
	}
}

func TestConnection_SetNodeA(t *testing.T) {
	n := Node{}
	c := Connection{}
	c.SetNodeA(&n)
	if c.GetNodeA() != &n {
		t.Fatal("Node A did not match expected Node A")
	}
}

func TestConnection_SetNodeB(t *testing.T) {
	n := Node{}
	c := Connection{}
	c.SetNodeB(&n)
	if c.GetNodeB() != &n {
		t.Fatal("Node B did not match expected Node B")
	}
}

func TestConnection_GetInnovationNumber(t *testing.T) {
	c := Connection{innovationNumber: 5}
	if c.GetInnovationNumber() != c.innovationNumber {
		t.Fatalf("Expected innovation number of 5, but got %v", c.GetInnovationNumber())
	}
}

func TestConnection_SetInnovationNumber(t *testing.T) {
	c := Connection{}
	c.SetInnovationNumber(5)
	if c.GetInnovationNumber() != 5 {
		t.Fatalf("Expected innovation number of 5, but got %v", c.GetInnovationNumber())
	}
}

func TestConnection_Clone(t *testing.T) {
	testConnection := Connection{weight: 5, innovationNumber: 5}
	copyConnection := testConnection.Clone()

	if testConnection.GetWeight() != copyConnection.GetWeight() {
		t.Fatalf("Expected test connection and copy connection to have the same weight, but test connection " +
			"has weight %v and copied connection has weight %v", testConnection.GetWeight(), copyConnection.GetWeight())
	}
	if testConnection.GetInnovationNumber() != copyConnection.GetInnovationNumber() {
		t.Fatalf("Expected test connection and copy connection to have the same innovation, but test " +
			"connection has innovation %v and copied connection has innovation %v",
			testConnection.GetInnovationNumber(), copyConnection.GetInnovationNumber())
	}
	if &testConnection == copyConnection {
		t.Fatalf("Test connection and copy connection have the same memory address %v and %v", &testConnection,
			copyConnection)
	}
}
