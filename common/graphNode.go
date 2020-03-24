package common

type GraphNode struct {
	location Location
	edges    []edge
}

func NewNode(location Location) Node {
	graphNode := GraphNode{}
	graphNode.location = location
	graphNode.edges = make([]edge, 0)
	return graphNode
}

type edge struct {
	from Node
	to   Node
}

func (g GraphNode) GetConnected() []Node {
	connected := make([]Node, len(g.edges))
	for _, e := range g.edges {
		connected = append(connected, e.to)
	}
	return connected
}

func (g GraphNode) Connect(that Node) bool {
	wasSuccessful := g.GetLocation() != that.GetLocation()
	for _, ed := range g.edges {
		if ed.to.GetLocation() == that.GetLocation() {
			wasSuccessful = false
		}
	}
	if wasSuccessful {
		g.edges = append(g.edges, edge{g, that})
		that.Connect(g)
	}
	return wasSuccessful
}

func (g GraphNode) Disconnect(that Node) bool {
	wasSuccessful := false
	for i, ed := range g.edges {
		if ed.to.GetLocation() == that.GetLocation() {
			g.edges = append(g.edges[:i], g.edges[i+1:]...)
			wasSuccessful = true
			that.Disconnect(g)
			break
		}
	}
	return wasSuccessful
}

func (g GraphNode) GetLocation() Location {
	return g.location
}

func (g GraphNode) Compare(that Node) bool {
	if that == nil {
		return false
	}
	return g.GetLocation().Compare(that.GetLocation())
}

func (g GraphNode) HardCompare(that Node) bool {
	if that == nil {
		return false
	}
	if !g.GetLocation().Compare(that.GetLocation()) {
		return false
	}
	thisConnected := g.GetConnected()
	thatConnected := that.GetConnected()
	if len(thisConnected) != len(thatConnected) {
		return false
	}
	// should have the same connected Nodes
	isEqual := true
	for _, thisConnectedNode := range thisConnected {
		equalFound := false
		for _, thatConnectedNode := range thatConnected {
			if thisConnectedNode.Compare(thatConnectedNode) {
				equalFound = true
			}
		}
		if !equalFound {
			isEqual = false
		}
	}
	return isEqual
}
