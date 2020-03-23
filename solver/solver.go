package solver

import . "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

type LabSolver interface {
	SolveLabyrinth(labyrinth Labyrinth, from Node, to Node) //TODO: Specify Solver Step Info
}
