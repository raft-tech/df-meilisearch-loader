package producer

import (
	meiliGo "github.com/meilisearch/meilisearch-go"
	"github.com/rs/zerolog/log"
	"meilisearch-loader/internal/meilisearch/configs"
	"meilisearch-loader/internal/model"
	"meilisearch-loader/internal/unmarshall"
	"strings"
	"time"
)

type MeilisearchProducer struct {
	Client          *meiliGo.Client
	BatchSize       int64
	Index           string
	IndexPrimaryKey string
}

func NewProducer(meiliCfg *configs.Config) MeilisearchProducer {
	return MeilisearchProducer{
		Client: meiliGo.NewClient(meiliGo.ClientConfig{
			Host:   meiliCfg.Url,
			APIKey: meiliCfg.ApiKey,
		}),
		BatchSize:       meiliCfg.BatchSize,
		Index:           meiliCfg.Index,
		IndexPrimaryKey: meiliCfg.IndexPrimaryKey,
	}
}

// PublishMessageBatch publishes a batch of records to Meilisearch.
func (p *MeilisearchProducer) PublishMessageBatch(msgChan <-chan model.Message) {
	var msgs []map[string]any
	var bufSize int64 = 0
	var publishedRecords int64 = 0
	for msg := range msgChan {
		var msgValueJson map[string]any
		if err := unmarshall.Into(&msgValueJson, strings.NewReader(string(msg.Value))); err != nil {
			log.Error().Msgf("Failed to unmarshal message value into json: %s", err)
		} else {
			msgs = append(msgs, msgValueJson)
			bufSize++
			if bufSize == p.BatchSize {
				log.Info().Msgf("Publishing batch of %d records to meilisearch...", bufSize)
				start := time.Now()
				if taskInfo, err := p.Client.Index(p.Index).AddDocuments(msgs, p.IndexPrimaryKey); err != nil {
					log.Error().Msgf("Error when publishing to Meilisearch: %s", err)
					continue
				} else {
					log.Info().Msg("Waiting for task to be completed...")
					if t, err := p.Client.WaitForTask(taskInfo.TaskUID); err != nil {
						log.Fatal().Msgf("Task failed to be processed: %s", err)
					} else {
						publishedRecords += bufSize
						log.Info().Msgf("Task complete: %+v", *t)
						log.Info().Msgf("Took: %s", time.Since(start))
						log.Info().Msgf("Published %d records so far...", publishedRecords)
						bufSize = 0
						msgs = []map[string]any{}
					}
				}
			}
		}
	}
}
