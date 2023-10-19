package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	kafkaConfig "meilisearch-loader/internal/kafka/config"
	"meilisearch-loader/internal/kafka/consumer"
	meiliConfig "meilisearch-loader/internal/meilisearch/config"
	"meilisearch-loader/internal/meilisearch/producer"
	"meilisearch-loader/internal/shared"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Starting up meilisearch loader...")

	meiliCfg := meiliConfig.NewConfig()
	meiliCfg.ParseEnvs()

	meiliProducer := producer.NewProducer(meiliCfg)
	log.Info().Msgf("Connecting to meilisearch at %s...", meiliCfg.Url)
	if h, err := meiliProducer.Client.Health(); err != nil {
		log.Error().Msgf("Meilisearch returned non-healthy response: %+v", h)
		log.Fatal().Msgf(err.Error())
	}

	kafkaCfg := kafkaConfig.NewConfig()
	kafkaCfg.ParseEnvs()

	kafkaConsumer := consumer.NewNoAuth(kafkaCfg.BrokerHost, kafkaCfg.SchemaRegUrl, kafkaCfg.Topic)
	defer kafkaConsumer.KafkaClient.Close()

	msgChan := make(chan shared.Message)
	go kafkaConsumer.DeserializeMessage(msgChan)
	go meiliProducer.PublishMessageBatch(msgChan)

	select {}
}