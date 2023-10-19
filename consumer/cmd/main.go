package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	kafkaConfig "meilisearch-loader/consumer/internal/kafka/config"
	"meilisearch-loader/consumer/internal/kafka/consumer"
	meiliConfig "meilisearch-loader/consumer/internal/meilisearch/config"
	"meilisearch-loader/consumer/internal/meilisearch/producer"
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

	kafkaConsumer.DoConsume(&meiliProducer)

	select {}
}
