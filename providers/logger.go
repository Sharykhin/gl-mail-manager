package providers

import (
	"os"

	"log"

	"github.com/Sharykhin/gl-mail-manager/contracts"
	"github.com/Sharykhin/gl-mail-manager/services/logger/file"
)

func Logger() contracts.Logger {
	if val, ok := container["logger"]; ok {
		log.Println("ok", ok)
		return val.(contracts.Logger)
	}
	dir, _ := os.Getwd()
	l, err := file.New(dir + "/logs")
	if err != nil {
		log.Fatalf("something went wrong:%v\n", err)
	}

	container["logger"] = l
	return l
}
