package common

import "math"

type GraphNode struct {
	location Location
	edges    []Node
}

func newNode(location Location) Node {
	graphNode := GraphNode{}
	graphNode.location = location
	graphNode.edges = make([]Node, 0)

	return graphNode
}

func (g GraphNode) getLocation() Location {
	return g.location
}

func (g GraphNode) isNeighbor(that Node) bool {
	if that == nil {
		return false
	}

	wasSuccessful := true
	thatX, thatY, thatZ := that.getLocation().As3DCoordinates()
	thisX, thisY, thisZ := g.getLocation().As3DCoordinates()
	dx := math.Abs(float64(int64(thatX) - int64(thisX)))
	dy := math.Abs(float64(int64(thatY) - int64(thisY)))
	dz := math.Abs(float64(int64(thatZ) - int64(thisZ)))
	isDxVarying := dx == gridStep
	isDyVarying := dy == gridStep
	isDzConstant := dz != gridStep
	isDxNotZero := dx != 0
	isDyNotZero := dy != 0
	isDzNotZero := dz != 0

	if isDxVarying {
		if isDyNotZero || isDzNotZero {
			wasSuccessful = false
		}
	} else {
		if isDyVarying {
			if isDxNotZero || isDzNotZero {
				wasSuccessful = false
			}
		} else {
			if isDzConstant || isDxNotZero || isDyNotZero {
				wasSuccessful = false
			}
		}
	}

	return wasSuccessful
}

func (g GraphNode) connect(that Node) (bool, Node, Node) {
	if that == nil {
		return false, g, that
	}

	wasSuccessful := g.isNeighbor(that)

	for _, ed := range g.edges {
		if ed.compare(that) {
			wasSuccessful = false
		}
	}

	if wasSuccessful {
		g.edges = append(g.edges, that)
		_, that, _ = that.connect(g)
		g.edges = append(g.edges[:len(g.edges)-1], that)
	}

	return wasSuccessful, g, that
}

func (g GraphNode) disconnect(that Node) (bool, Node, Node) {
	if that == nil {
		return false, g, that
	}

	wasSuccessful := false

	for i, ed := range g.edges {
		if ed.compare(that) {
			wasSuccessful = true

			g.edges = append(g.edges[:i], g.edges[i+1:]...)

			_, that, _ = that.disconnect(g)
		}
	}

	return wasSuccessful, g, that
}

func (g GraphNode) getConnected() []Node {
	connected := make([]Node, 0)

	connected = append(connected, g.edges...)

	return connected
}

func (g GraphNode) compare(that Node) bool {
	if that == nil {
		return false
	}

	return g.getLocation().Compare(that.getLocation())
}

func (g GraphNode) hardCompare(that Node) bool {
	if !g.compare(that) {
		return false
	}

	thisConnected := g.getConnected()
	thatConnected := that.getConnected()

	if len(thisConnected) != len(thatConnected) {
		return false
	}

	// should have the same connected Nodes
	isEqual := true

	for _, thisConnectedNode := range thisConnected {
		equalFound := false

		for _, thatConnectedNode := range thatConnected {
			if (thisConnectedNode).compare(thatConnectedNode) {
				equalFound = true
			}
		}

		if !equalFound {
			isEqual = false
		}
	}

	return isEqual
}
