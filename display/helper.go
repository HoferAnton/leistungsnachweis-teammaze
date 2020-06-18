package display

import (
	"log"

	"github.com/go-gl/gl/v4.2-core/gl"
)

func FatalIfError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func checkForGLError(message string) {
	err := gl.GetError()
	if err != gl.NO_ERROR {
		log.Fatalf("%s ErrorCode: 0x%X", message, err)
	}
}
