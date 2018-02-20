package file

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	log *log.Logger
}

const ERRORFILENAME = "error.log"

func New(dirPath string) (*Logger, error) {
	var l = new(Logger)
	var filename string = dirPath + "/" + ERRORFILENAME
	var file, err = os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("could not create file:%s. %v", filename, err)
	}
	l.log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	l.log.SetFlags(1)
	return l, nil
}

func (l Logger) LogError(str string) {
	l.log.Println(str)
}
