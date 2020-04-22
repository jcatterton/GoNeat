package Network

type Connection struct {
	weight				float64
	nodeA				*Node
	nodeB				*Node
	innovationNumber	int
}

func (c *Connection) GetWeight() float64 {
	return c.weight
}

func (c *Connection) SetWeight(f float64) {
	c.weight = f
}

func (c *Connection) GetNodeA() *Node {
	return c.nodeA
}

func (c *Connection) GetNodeB() *Node {
	return c.nodeB
}

func (c *Connection) SetNodeA(n *Node) {
	c.nodeA = n
}

func (c *Connection) SetNodeB(n *Node) {
	c.nodeB = n
}

func (c *Connection) GetInnovationNumber() int {
	return c.innovationNumber
}

func (c *Connection) SetInnovationNumber(i int) {
	c.innovationNumber = i
}

func (c *Connection) Clone() *Connection {
	newConnection := &Connection{weight: c.GetWeight(), innovationNumber: c.GetInnovationNumber()}
	return newConnection
}
