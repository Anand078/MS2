package main

import (
	appKafka "MS2/kafka"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := appKafka.StartKafka()
	if err != nil {
		log.Infoln("There is an error while reading data from kafka..", err)
	}
}
