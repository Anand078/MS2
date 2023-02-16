package main

import (
	appKafka "MS2/kafkaimp"
	"fmt"
	"time"
)

func main() {
	go appKafka.StartKafka()

	fmt.Println("Kafka has been started...")

	time.Sleep(10 * time.Minute)
}
