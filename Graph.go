package main

type GraphNode struct {
	Location Location
	edges    []edge
}

type edge struct {
	from Node
	to Node
}

func New(location Location) GraphNode {
	graphNode := GraphNode{}
	graphNode.Location = location
	graphNode.edges = make([]edge, 0)
	return graphNode

}


func (g GraphNode) GetNeighbors() []Node {
	neighbors := make([]Node, len(g.edges))
	for e := range g.edges {
		neighbors = append(neighbors, e.to)
	}
	return
}

func (g GraphNode) GetConnected() []Node {
	panic("implement me")
}

func (g GraphNode) Connect(Node) bool {
	panic("implement me")
}

func (g GraphNode) Disconnect(Node) bool {
	panic("implement me")
}

func (g GraphNode) GetLocation() Location {
	panic("implement me")
}

func (g GraphNode) Equals(Node) bool {
	panic("implement me")
}

