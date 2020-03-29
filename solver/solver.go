package solver

import "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

type LabSolver interface {
	SolveLabyrinth(labyrinth common.Labyrinth, from common.Node, to common.Node)
}
