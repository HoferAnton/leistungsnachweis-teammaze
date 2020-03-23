package generator

import . "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

type LabGenerator interface {
	GenerateLabyrinth(furthestPoint Location) Labyrinth //TODO: Specify Generation Step Info
}
