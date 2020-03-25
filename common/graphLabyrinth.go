package common

type GraphLabyrinth struct {
	nodes  []Node
	maxLoc Location
}

func NewLabyrinth(maxLoc Location) Labyrinth {
	if maxLoc == nil {
		return nil
	}
	graphNode := GraphLabyrinth{}
	graphNode.nodes = make([]Node, 0)
	maxX, maxY, maxZ := maxLoc.As3DCoordinates()
	for z := uint(0); z <= maxZ; z++ {
		for y := uint(0); y <= maxY; y++ {
			for x := uint(0); x <= maxX; x++ {
				graphNode.nodes = append(graphNode.nodes, newNode(NewLocation(x, y, z)))
			}
		}
	}
	graphNode.maxLoc = maxLoc
	return graphNode
}

func (g GraphLabyrinth) GetMaxLocation() Location {
	return g.maxLoc
}

func (g GraphLabyrinth) GetNeighbors(loc Location) []Location {
	if loc == nil {
		return nil
	}
	x, y, z := loc.As3DCoordinates()
	maxX, maxY, maxZ := g.GetMaxLocation().As3DCoordinates()
	if !g.checkLocation(loc) {
		return nil
	}

	neighbors := make([]Location, 0)
	if x != 0 {
		neighbors = append(neighbors, g.nodes[getIndex(x-1, y, z, g.maxLoc)].getLocation())
	}
	if x != maxX {
		neighbors = append(neighbors, g.nodes[getIndex(x+1, y, z, g.maxLoc)].getLocation())
	}
	if y != 0 {
		neighbors = append(neighbors, g.nodes[getIndex(x, y-1, z, g.maxLoc)].getLocation())
	}
	if y != maxY {
		neighbors = append(neighbors, g.nodes[getIndex(x, y+1, z, g.maxLoc)].getLocation())
	}
	if z != 0 {
		neighbors = append(neighbors, g.nodes[getIndex(x, y, z-1, g.maxLoc)].getLocation())
	}
	if z != maxZ {
		neighbors = append(neighbors, g.nodes[getIndex(x, y, z+1, g.maxLoc)].getLocation())
	}
	return neighbors
}

func (g GraphLabyrinth) Connect(loc1 Location, loc2 Location) bool {
	node1 := g.getNode(loc1)
	node2 := g.getNode(loc2)
	if node1 == nil || node2 == nil ||
		!g.checkLocation(loc1) ||
		!g.checkLocation(loc1) {
		return false
	}
	wasSuccessful, node1, node2 := node1.connect(node2)

	g.nodes = replaceNodes(node1, g.nodes, g.GetMaxLocation())
	g.nodes = replaceNodes(node2, g.nodes, g.GetMaxLocation())

	return wasSuccessful
}

func (g GraphLabyrinth) Disconnect(loc1 Location, loc2 Location) bool {
	node1 := g.getNode(loc1)
	node2 := g.getNode(loc2)
	if node1 == nil || node2 == nil ||
		!g.checkLocation(loc1) ||
		!g.checkLocation(loc1) {
		return false
	}
	wasSuccessful, node1, node2 := node1.disconnect(node2)

	g.nodes = replaceNodes(node1, g.nodes, g.GetMaxLocation())
	g.nodes = replaceNodes(node2, g.nodes, g.GetMaxLocation())

	return wasSuccessful
}

func (g GraphLabyrinth) GetConnected(loc Location) []Location {
	node := g.getNode(loc)
	if node == nil || !g.checkLocation(loc) {
		return nil
	}
	connected := node.getConnected()
	conLoc := make([]Location, 0)
	for _, con := range connected {
		conLoc = append(conLoc, con.getLocation())
	}
	return conLoc
}

func (g GraphLabyrinth) Compare(that Labyrinth) bool {
	if that == nil {
		return false
	}
	// all nodes should be equal
	for _, n := range g.nodes {
		thatNode := that.getNode(n.getLocation())
		if thatNode == nil || !thatNode.hardCompare(n) {
			return false
		}
	}
	// both should not have further nodes
	maxX, maxY, maxZ := g.maxLoc.As3DCoordinates()
	locX := NewLocation(maxX+1, maxY, maxZ)
	lastTestNodeX := g.getNode(locX)
	thatNodeX := that.getNode(locX)
	locY := NewLocation(maxX, maxY+1, maxZ)
	lastTestNodeY := g.getNode(locY)
	thatNodeY := that.getNode(locY)
	locZ := NewLocation(maxX, maxY, maxZ+1)
	lastTestNodeZ := g.getNode(locZ)
	thatNodeZ := that.getNode(locZ)
	return thatNodeX == nil &&
		thatNodeX == lastTestNodeX &&
		thatNodeY == nil &&
		thatNodeY == lastTestNodeY &&
		thatNodeZ == nil &&
		thatNodeZ == lastTestNodeZ
}

func (g GraphLabyrinth) checkLocation(loc Location) bool {
	if loc == nil {
		return false
	}
	x, y, z := loc.As3DCoordinates()
	maxX, maxY, maxZ := g.GetMaxLocation().As3DCoordinates()
	return x <= maxX && y <= maxY && z <= maxZ
}

func (g GraphLabyrinth) getNode(location Location) Node {
	if location == nil {
		return nil
	}
	maxX, maxY, maxZ := g.maxLoc.As3DCoordinates()
	x, y, z := location.As3DCoordinates()
	if x <= maxX && y <= maxY && z <= maxZ {
		return g.nodes[getIndex(x, y, z, g.maxLoc)]
	} else {
		return nil
	}
}

func replaceNodes(node Node, nodes []Node, maxLoc Location) []Node {
	x, y, z := node.getLocation().As3DCoordinates()
	index := getIndex(x, y, z, maxLoc)
	tmp := nodes[index+1:]
	nodes = append(nodes[:index], node)
	nodes = append(nodes, tmp...)
	return nodes
}

func getIndex(x uint, y uint, z uint, maxLoc Location) uint {
	maxX, maxY, _ := maxLoc.As3DCoordinates()
	return x + y*(maxX+1) + z*(maxX+1)*(maxY+1)
}
