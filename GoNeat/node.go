package GoNeat

import (
	"math"
)

type Node struct {
	activated          bool
	layer              int
	outwardConnections []*Connection
	inwardConnections  []*Connection
	weight             float64
	innovationNumber   int
}

func CreateNode(activated bool, layer int, outwardConnections []*Connection, inwardConnections []*Connection, weight float64, innovation int) *Node {
	return &Node{activated, layer, outwardConnections, inwardConnections, weight, innovation}
}

func (n *Node) IsActivated() bool {
	return n.activated
}

func (n *Node) Activate() {
	n.activated = true
}

func (n *Node) Deactivate() {
	n.activated = false
}

func (n *Node) GetOutwardConnections() []*Connection {
	return n.outwardConnections
}

func (n *Node) GetInwardConnections() []*Connection {
	return n.inwardConnections
}

func (n *Node) AddToOutwardConnections(c *Connection) {
	n.outwardConnections = append(n.outwardConnections, c)
}

func (n *Node) AddToInwardConnections(c *Connection) {
	n.inwardConnections = append(n.inwardConnections, c)
}

func (n *Node) RemoveFromOutwardConnections(c *Connection) int {
	for i := range n.outwardConnections {
		if n.outwardConnections[i] == c {
			n.outwardConnections[len(n.outwardConnections)-1], n.outwardConnections[i] =
				n.outwardConnections[i], n.outwardConnections[len(n.outwardConnections)-1]
			n.outwardConnections = n.outwardConnections[:len(n.outwardConnections)-1]
			return 1
		}
	}
	return -1
}

func (n *Node) RemoveFromInwardConnections(c *Connection) int {
	for i := range n.inwardConnections {
		if n.inwardConnections[i] == c {
			n.inwardConnections[len(n.inwardConnections)-1], n.inwardConnections[i] =
				n.inwardConnections[i], n.inwardConnections[len(n.inwardConnections)-1]
			n.inwardConnections = n.inwardConnections[:len(n.inwardConnections)-1]
			return 1
		}
	}
	return -1
}

func (n *Node) ClearOutwardConnections() {
	n.outwardConnections = []*Connection{}
}

func (n *Node) ClearInwardConnections() {
	n.inwardConnections = []*Connection{}
}

func (n *Node) GetWeight() float64 {
	return n.weight
}

func (n *Node) SetWeight(w float64) {
	n.weight = w
}

func (n *Node) GetInnovationNumber() int {
	return n.innovationNumber
}

func (n *Node) SetInnovationNumber(i int) {
	n.innovationNumber = i
}

func (n *Node) Sigmoid() {
	n.weight = 1 / (1 + math.Exp(n.GetWeight()*-4.9))
	if n.weight <= 0.5 {
		n.Deactivate()
	} else {
		n.Activate()
	}
}

func (n *Node) GetLayer() int {
	return n.layer
}

func (n *Node) SetLayer(l int) {
	n.layer = l
}

func (n *Node) IsConnectedTo(o *Node) bool {
	for i := range n.GetOutwardConnections() {
		if n.GetOutwardConnections()[i].GetNodeB() == o {
			return true
		}
	}
	for i := range n.GetInwardConnections() {
		if n.GetInwardConnections()[i].GetNodeA() == o {
			return true
		}
	}
	return false
}

func (n *Node) Clone() *Node {
	newNode := &Node{activated: n.IsActivated(), layer: n.GetLayer(), weight: n.GetWeight(),
		innovationNumber: n.GetInnovationNumber()}
	return newNode
}
