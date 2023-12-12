package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	kafkaConfig "meilisearch-loader/internal/kafka/configs"
	"meilisearch-loader/internal/kafka/consumer"
	meiliConfig "meilisearch-loader/internal/meilisearch/configs"
	"meilisearch-loader/internal/meilisearch/producer"
	"meilisearch-loader/internal/model"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Starting up meilisearch loader...")

	meiliCfg := meiliConfig.NewConfig()

	meiliProducer := producer.NewProducer(meiliCfg)
	log.Info().Msgf("Connecting to meilisearch at %s...", meiliCfg.Url)
	if h, err := meiliProducer.Client.Health(); err != nil {
		log.Error().Msgf("Meilisearch returned non-healthy response: %+v", h)
		log.Fatal().Msgf(err.Error())
	}

	kafkaCfg := kafkaConfig.NewConfig()

	var kafkaConsumer consumer.DeserializingAvroConsumer
	if kafkaCfg.SaslMechanism == nil {
		kafkaConsumer = consumer.NewNoAuth(kafkaCfg.BrokerHost, kafkaCfg.SchemaRegUrl, kafkaCfg.Topic)
	} else {
		kafkaConsumer = consumer.NewSaslAuth(kafkaCfg.BrokerHost, kafkaCfg.SchemaRegUrl, kafkaCfg.Topic, kafkaCfg.SaslMechanism)
	}
	defer kafkaConsumer.KafkaClient.Close()

	msgChan := make(chan model.Message)
	go kafkaConsumer.DeserializeMessage(msgChan)
	go meiliProducer.PublishMessageBatch(msgChan)

	select {}
}
