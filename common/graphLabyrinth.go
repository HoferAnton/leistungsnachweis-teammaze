package common

type GraphLabyrinth struct {
	nodes  []Node
	maxLoc Location
}

func NewLabyrinth(maxLoc Location) Labyrinth {
	graphNode := GraphLabyrinth{}
	graphNode.nodes = make([]Node, 0)
	maxX, maxY, maxZ := maxLoc.As3DCoordinates()
	for z := uint(0); z <= maxZ; z++ {
		for y := uint(0); y <= maxY; y++ {
			for x := uint(0); x <= maxX; x++ {
				graphNode.nodes = append(graphNode.nodes, NewNode(NewLocation(x, y, z)))
			}
		}
	}
	graphNode.maxLoc = maxLoc
	return graphNode
}

func (g GraphLabyrinth) GetNode(location Location) Node {
	maxX, maxY, maxZ := g.maxLoc.As3DCoordinates()
	x, y, z := location.As3DCoordinates()
	if x <= maxX && y <= maxY && z <= maxZ {
		return g.nodes[getIndex(x, y, z, g.maxLoc)]
	} else {
		return nil
	}
}

func (g GraphLabyrinth) GetNeighborsByNode(node Node) []Node {
	return g.GetNeighborsByLocation(node.GetLocation())
}

func (g GraphLabyrinth) GetNeighborsByLocation(loc Location) []Node {
	x, y, z := loc.As3DCoordinates()
	maxX, maxY, maxZ := g.maxLoc.As3DCoordinates()
	if x > maxX || y > maxY || z > maxZ {
		return nil
	}

	neighbors := make([]Node, 0)
	if x != 0 {
		neighbors = append(neighbors, g.nodes[getIndex(x-1, y, z, g.maxLoc)])
	}
	if x != maxX {
		neighbors = append(neighbors, g.nodes[getIndex(x+1, y, z, g.maxLoc)])
	}
	if y != 0 {
		neighbors = append(neighbors, g.nodes[getIndex(x, y-1, z, g.maxLoc)])
	}
	if y != maxY {
		neighbors = append(neighbors, g.nodes[getIndex(x, y+1, z, g.maxLoc)])
	}
	if z != 0 {
		neighbors = append(neighbors, g.nodes[getIndex(x, y, z-1, g.maxLoc)])
	}
	if z != maxZ {
		neighbors = append(neighbors, g.nodes[getIndex(x, y, z+1, g.maxLoc)])
	}
	return neighbors
}

func (g GraphLabyrinth) Compare(that Labyrinth) bool {
	if that == nil {
		return false
	}
	// all nodes should be equal
	for _, n := range g.nodes {
		thatNode := that.GetNode(n.GetLocation())
		if thatNode == nil || !thatNode.HardCompare(n) {
			return false
		}
	}
	// both should not have further nodes
	maxX, maxY, maxZ := g.maxLoc.As3DCoordinates()
	locX := NewLocation(maxX+1, maxY, maxZ)
	lastTestNodeX := g.GetNode(locX)
	thatNodeX := that.GetNode(locX)
	locY := NewLocation(maxX, maxY+1, maxZ)
	lastTestNodeY := g.GetNode(locY)
	thatNodeY := that.GetNode(locY)
	locZ := NewLocation(maxX, maxY, maxZ+1)
	lastTestNodeZ := g.GetNode(locZ)
	thatNodeZ := that.GetNode(locZ)
	return thatNodeX == nil &&
		thatNodeX == lastTestNodeX &&
		thatNodeY == nil &&
		thatNodeY == lastTestNodeY &&
		thatNodeZ == nil &&
		thatNodeZ == lastTestNodeZ
}

func getIndex(x uint, y uint, z uint, maxLoc Location) uint {
	maxX, maxY, _ := maxLoc.As3DCoordinates()
	return x + y*(maxX+1) + z*(maxX+1)*(maxY+1)
}
