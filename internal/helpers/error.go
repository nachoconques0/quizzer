package helpers

import (
	"log"
)

func HandleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%v err: %v", msg, err)
	}
}
