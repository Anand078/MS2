package kafkaimp

import (
	"context"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func StartKafka() error {
	log.Infoln("StartKafka is started...")
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test",
		GroupID:  "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Infoln("some error occurred", err)
			return err
		}
		log.Infoln("Message is : ", string(m.Value))
	}
}
