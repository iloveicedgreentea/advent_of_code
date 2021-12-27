package common

import (
	log "github.com/sirupsen/logrus"
)

// checks an error 
func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}