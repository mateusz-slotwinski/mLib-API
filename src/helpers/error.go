package helpers

import (
	context "context"
	log "log"
	Config "mLibAPI/src/config"
)

func PrintError(err error) {
	if Config.Mode == "DEV" && err != nil {
		log.Fatal(err)
	}
}

func PrintCancel(err context.CancelFunc) {
	if Config.Mode == "DEV" && err != nil {
		log.Fatal(err)
	}
}
