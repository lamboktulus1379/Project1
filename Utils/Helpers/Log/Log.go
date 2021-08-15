package Log

import (
	"log"
)

func INFO(msg string, args ...string) {

	log.Printf("INFO: "+msg, args[0], args[1])
}

func ERROR(msg string, args ...string) {
	log.Printf("ERROR: " + msg)
}

func DEBUG(msg string, args ...string) {
	log.Printf("DEBUG: " + msg)
}
