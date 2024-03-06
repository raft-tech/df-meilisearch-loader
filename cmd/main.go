package main

import (
	"meilisearch-loader/internal/kafka/consumer"
	"meilisearch-loader/internal/meilisearch/producer"
	"meilisearch-loader/internal/model"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	kafkaConfig "meilisearch-loader/internal/kafka/configs"

	meiliConfig "meilisearch-loader/internal/meilisearch/configs"
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

	kafkaCfg := kafkaConfig.NewConfig(meiliCfg.Index)

	var kafkaConsumer consumer.DeserializingAvroConsumer
	if kafkaCfg.SaslMechanism == "" {
		kafkaConsumer = consumer.NewNoAuth(kafkaCfg.BrokerHost, kafkaCfg.SchemaRegUrl, kafkaCfg.Topic, kafkaCfg.GroupId)
	} else {
		kafkaConsumer = consumer.NewSaslAuth(kafkaCfg.BrokerHost, kafkaCfg.SchemaRegUrl, kafkaCfg.Topic, kafkaCfg.GroupId, kafkaCfg.SaslMechanism, kafkaCfg.SaslUsername, kafkaCfg.SaslSecret)
	}
	defer kafkaConsumer.KafkaClient.Close()

	msgChan := make(chan model.Message)
	go kafkaConsumer.DeserializeMessage(msgChan)
	go meiliProducer.PublishMessageBatch(msgChan)

	select {}
}
