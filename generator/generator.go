package generator

import "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

type LabGenerator interface {
	GenerateLabyrinth(furthestPoint common.Location) common.Labyrinth //TODO: Specify Generation Step Info
}
