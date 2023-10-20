# Meilisearch Loader

Application to load data from a Kafka topic and push the data to Meilisearch

### ENVIRONMENT VARIABLES

The following table lists the environment variables that can be set, along with their default values.

| Name                                                                                                               | Description                                                 | Default          |
|--------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|------------------|
| KAFKA_BROKER_HOST                                                                                                  | The broker address used to connect to the kafka cluster     | localhost:9092   |
| KAFKA_TOPIC                                                                                                        | The topic to consume the data from                          | test-topic       |
| [SCHEMA_REGISTRY_URL](https://docs.confluent.io/platform/current/schema-registry/index.html)                       | The URL of the schema registry                              | localhost:8081   |
| MEILISEARCH_URL                                                                                                    | The address of the server Meilisearch is listening on       | localhost:7700   |
| [MEILISEARCH_INDEX](https://www.meilisearch.com/docs/learn/core_concepts/indexes)                                  | The uid where documents will be grouped                     | test-index       |
| [MEILISEARCH_INSERT_BATCH_SIZE](https://www.meilisearch.com/docs/reference/api/documents#add-or-replace-documents) | The number of documents to batch before adding to the index | 50               |
| [MEILISEARCH_API_KEY](https://www.meilisearch.com/docs/learn/security/master_api_keys)                             | Protects against unauthorized requests                      | aSampleMasterKey |
| [MEILISEARCH_INDEX_PRIMARY_KEY](https://www.meilisearch.com/docs/learn/core_concepts/primary_key)                  | Uniquely identifies each document in an index               | id               |
