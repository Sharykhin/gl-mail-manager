package logger

import (
	"log"
	"os"
)

// Log is a reference to a private struct
var Log logger

type logger struct {
	log *log.Logger
}

func (l logger) LogError(str string) {
	l.log.Println(str)
}

func init() {
	var filename = "logs/error.log"
	var file, err = os.Create(filename)
	if err != nil {
		log.Fatalf("Could not create a file %s: %v \n", filename, err)
	}
	Log.log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.log.SetFlags(1)
}
