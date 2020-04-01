package common

type GraphLabyrinth struct {
	nodes  []Node
	maxLoc Location
}

func NewLabyrinth(maxLoc Location) Labyrinth {
	if maxLoc == nil {
		return nil
	}

	graphLabyrinth := GraphLabyrinth{}
	graphLabyrinth.nodes = make([]Node, 0)
	maxX, maxY, maxZ := maxLoc.As3DCoordinates()

	for z := uint(0); z <= maxZ; z++ {
		for y := uint(0); y <= maxY; y++ {
			for x := uint(0); x <= maxX; x++ {
				graphLabyrinth.nodes = append(graphLabyrinth.nodes, newNode(NewLocation(x, y, z)))
			}
		}
	}

	graphLabyrinth.maxLoc = maxLoc

	return graphLabyrinth
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

	if !g.CheckLocation(loc) {
		return nil
	}

	neighbors := make([]Location, 0)

	if x != 0 {
		neighbors = append(neighbors, g.nodes[GetIndex(x-gridStep, y, z, g.maxLoc)].getLocation())
	}

	if x != maxX {
		neighbors = append(neighbors, g.nodes[GetIndex(x+gridStep, y, z, g.maxLoc)].getLocation())
	}

	if y != 0 {
		neighbors = append(neighbors, g.nodes[GetIndex(x, y-gridStep, z, g.maxLoc)].getLocation())
	}

	if y != maxY {
		neighbors = append(neighbors, g.nodes[GetIndex(x, y+gridStep, z, g.maxLoc)].getLocation())
	}

	if z != 0 {
		neighbors = append(neighbors, g.nodes[GetIndex(x, y, z-gridStep, g.maxLoc)].getLocation())
	}

	if z != maxZ {
		neighbors = append(neighbors, g.nodes[GetIndex(x, y, z+gridStep, g.maxLoc)].getLocation())
	}

	return neighbors
}

func (g GraphLabyrinth) Connect(loc1 Location, loc2 Location) bool {
	node1 := g.getNode(loc1)
	node2 := g.getNode(loc2)

	if node1 == nil || node2 == nil ||
		!g.CheckLocation(loc1) ||
		!g.CheckLocation(loc1) {
		return false
	}

	wasSuccessful, node1, node2 := node1.connect(node2)
	x1, y1, z1 := node1.getLocation().As3DCoordinates()
	g.nodes[GetIndex(x1, y1, z1, g.GetMaxLocation())] = node1
	x2, y2, z2 := node2.getLocation().As3DCoordinates()
	g.nodes[GetIndex(x2, y2, z2, g.GetMaxLocation())] = node2

	return wasSuccessful
}

func (g GraphLabyrinth) Disconnect(loc1 Location, loc2 Location) bool {
	node1 := g.getNode(loc1)
	node2 := g.getNode(loc2)

	if node1 == nil || node2 == nil ||
		!g.CheckLocation(loc1) ||
		!g.CheckLocation(loc1) {
		return false
	}

	wasSuccessful, node1, node2 := node1.disconnect(node2)
	x1, y1, z1 := node1.getLocation().As3DCoordinates()
	g.nodes[GetIndex(x1, y1, z1, g.GetMaxLocation())] = node1
	x2, y2, z2 := node2.getLocation().As3DCoordinates()
	g.nodes[GetIndex(x2, y2, z2, g.GetMaxLocation())] = node2

	return wasSuccessful
}

func (g GraphLabyrinth) GetConnected(loc Location) []Location {
	node := g.getNode(loc)

	if node == nil || !g.CheckLocation(loc) {
		return nil
	}

	connected := node.getConnected()
	conLoc := make([]Location, 0)

	for _, con := range connected {
		conLoc = append(conLoc, con.getLocation())
	}

	return conLoc
}

func (g GraphLabyrinth) IsConnected(loc1 Location, loc2 Location) bool {
	node1 := g.getNode(loc1)
	node2 := g.getNode(loc2)

	if node1 == nil || node2 == nil || !g.CheckLocation(loc1) || !g.CheckLocation(loc2) {
		return false
	}

	connected := node1.getConnected()

	for _, con := range connected {
		if con.compare(node2) {
			return true
		}
	}

	return false
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
	locX := NewLocation(maxX+gridStep, maxY, maxZ)
	lastTestNodeX := g.getNode(locX)
	thatNodeX := that.getNode(locX)
	locY := NewLocation(maxX, maxY+gridStep, maxZ)
	lastTestNodeY := g.getNode(locY)
	thatNodeY := that.getNode(locY)
	locZ := NewLocation(maxX, maxY, maxZ+gridStep)
	lastTestNodeZ := g.getNode(locZ)
	thatNodeZ := that.getNode(locZ)

	return thatNodeX == nil &&
		thatNodeX == lastTestNodeX &&
		thatNodeY == nil &&
		thatNodeY == lastTestNodeY &&
		thatNodeZ == nil &&
		thatNodeZ == lastTestNodeZ
}

func (g GraphLabyrinth) CheckLocation(loc Location) bool {
	if loc == nil {
		return false
	}

	x, y, z := loc.As3DCoordinates()
	maxX, maxY, maxZ := g.GetMaxLocation().As3DCoordinates()

	return x <= maxX && y <= maxY && z <= maxZ
}

func (g GraphLabyrinth) getNode(location Location) Node {
	if g.CheckLocation(location) {
		x, y, z := location.As3DCoordinates()
		return g.nodes[GetIndex(x, y, z, g.maxLoc)]
	}

	return nil
}

func GetIndex(x uint, y uint, z uint, maxLoc Location) uint {
	maxX, maxY, _ := maxLoc.As3DCoordinates()
	return x + y*(maxX+1) + z*(maxX+1)*(maxY+1)
}
