package main

type Location3D struct {
	x, y, z int
}

type Location interface {
	As3DCoordinates() (int, int, int)
}

type Labyrinth struct {
	nodes []Node
}

type LabGenerator interface {
	GenerateLabyrinth(furthestPoint Location) Labyrinth //TODO: Specify Generation Step Info
}

type LabSolver interface {
	SolveLabyrinth(labyrinth Labyrinth, from Node, to Node) //TODO: Specify Solver Step Info
}

type Node interface {
	GetNeighbors() []Node
	GetConnected() []Node

	Connect(Node) bool
	Disconnect(Node) bool

	GetLocation() Location

	Equals(Node) bool
}