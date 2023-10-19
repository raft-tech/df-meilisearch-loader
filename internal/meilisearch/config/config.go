package config

import (
	"os"
	"strconv"
)

type Config struct {
	ApiKey          string
	BatchSize       int64
	Index           string
	IndexPrimaryKey string
	Url             string
}

var (
	meilisearchUrl                     = "localhost:7700"
	meilisearchApiKey                  = "aSampleMasterKey"
	meilisearchIndex                   = "test-index"
	meilisearchBatchSize         int64 = 50
	meilisearchDatasetPrimaryKey       = "id"
)

func NewConfig() *Config {
	return &Config{
		ApiKey:          meilisearchApiKey,
		BatchSize:       meilisearchBatchSize,
		IndexPrimaryKey: meilisearchDatasetPrimaryKey,
		Index:           meilisearchIndex,
		Url:             meilisearchUrl,
	}
}

func (cfg *Config) ParseEnvs() {
	if m, exists := os.LookupEnv("MEILISEARCH_URL"); exists {
		cfg.Url = m
	}
	if mKey, exists := os.LookupEnv("MEILISEARCH_API_KEY"); exists {
		cfg.ApiKey = mKey
	}
	if mIndex, exists := os.LookupEnv("MEILISEARCH_INDEX"); exists {
		cfg.Index = mIndex
	}
	if mBatchSize, exists := os.LookupEnv("MEILISEARCH_INSERT_BATCH_SIZE"); exists {
		if tryMBatchSize, err := strconv.ParseInt(mBatchSize, 10, 64); err == nil {
			cfg.BatchSize = tryMBatchSize
		}
	}
	if mKey, exists := os.LookupEnv("MEILISEARCH_INDEX_PRIMARY_KEY"); exists {
		cfg.IndexPrimaryKey = mKey
	}
}
