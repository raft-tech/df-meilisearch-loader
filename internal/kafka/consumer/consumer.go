package consumer

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/rs/zerolog/log"
	"meilisearch-loader/internal/model"
	"meilisearch-loader/internal/unmarshall"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/linkedin/goavro/v2"
	"github.com/segmentio/kafka-go"
)

const AvroMessageMagicByte = 0x0

type SchemaRegistryResponse struct {
	Subject string `json:"subject,omitempty"`
	Version int    `json:"version,omitempty"`
	Id      uint32 `json:"id,omitempty"`
	Schema  string `json:"schema"`
}

type DeserializingAvroConsumer struct {
	BrokerHost         string
	KafkaClient        *kafka.Reader
	SchemaRegistryHost string
	Topic              string
	KeySchema          *Schema
	ValueSchema        *Schema
}

type Schema struct {
	Raw string
	// Version int // Could set this but isn't used
	Id    uint32
	Codec *goavro.Codec
}

func NewNoAuth(kafkaHost, schemaRegHost, topic string) DeserializingAvroConsumer {
	// mechanism := plain.Mechanism{Username: saslUser, Password: saslPass}
	return DeserializingAvroConsumer{
		BrokerHost:         kafkaHost,
		SchemaRegistryHost: schemaRegHost,
		KafkaClient: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{kafkaHost},
			Topic:   topic,
			Dialer:  &kafka.Dialer{Timeout: 10 * time.Second, DualStack: true, SASLMechanism: nil},
		}),
		Topic:       topic,
		KeySchema:   nil,
		ValueSchema: nil,
	}
}

// DeserializeMessage deserializes the key and value from a message and adds the deserialized message to msgChan
func (c *DeserializingAvroConsumer) DeserializeMessage(msgChan chan<- model.Message) {
	for {
		message := model.Message{}
		m, err := c.KafkaClient.ReadMessage(context.Background())
		if err != nil {
			log.Error().Msgf("Error reading from Kafka: %s", err)
			close(msgChan)
			return
		}
		if len(m.Key) > 0 {
			deserializedMsg, err := c.deserializeKey(&m)
			if err != nil {
				log.Error().Msgf("Error attempting to deserialize key: %s", err)
				close(msgChan)
				return
			}
			message.Key = deserializedMsg
		}
		if len(m.Value) > 0 {
			deserializedMsg, err := c.deserializeValue(&m)
			if err != nil {
				log.Error().Msgf("Error attempting to deserialize value: %s", err)
				close(msgChan)
				return
			}
			message.Value = deserializedMsg
		}
		msgChan <- message
	}
}

func (c *DeserializingAvroConsumer) deserializeKey(m *kafka.Message) ([]byte, error) {
	return c.deserializeAvro(m, true)
}

func (c *DeserializingAvroConsumer) deserializeValue(m *kafka.Message) ([]byte, error) {
	return c.deserializeAvro(m, false)
}

// deserializeAvro deserializes a message if it has an Avro format, otherwise leaves the message as is
func (c *DeserializingAvroConsumer) deserializeAvro(msg *kafka.Message, deserKey bool) ([]byte, error) {
	var err error
	var out []byte
	var avroBytes []byte
	var schema *Schema
	if deserKey {
		avroBytes = msg.Key
		schema = c.KeySchema
	} else {
		avroBytes = msg.Value
		schema = c.ValueSchema
	}
	// Check magic byte is 0x0 (else not an Avro message)
	if avroBytes[0] != AvroMessageMagicByte {
		log.Info().Msgf("Not an Avro Message; invalid magic byte: (%b) != %b", avroBytes[0], AvroMessageMagicByte)
		log.Info().Msg("Skipping Avro deserialization.")
		out = avroBytes
	} else {
		// Skip over the first 5 bytes (magic byte + 4 bytes for schema ID)
		// Retrieve the Schema ID from the message
		msgSchemaId := binary.BigEndian.Uint32(avroBytes[1:5])
		// If no schema is present or its ID is different from the message schema ID, attempt to retrieve the new one
		if schema == nil || msgSchemaId != schema.Id {
			var existingSchemaId uint32 = 0
			if deserKey && schema != nil {
				existingSchemaId = c.KeySchema.Id
			} else if !deserKey && schema != nil {
				existingSchemaId = c.ValueSchema.Id
			}
			log.Info().Msgf("##### Detected change in schema. New schema ID: %d (previously %d)", msgSchemaId, existingSchemaId)
			schema, err = retrieveSchemaById(c.SchemaRegistryHost, msgSchemaId)
			if err != nil {
				return out, fmt.Errorf("ERROR: Failed to retrieve new schema: %s", err)
			} else if schema == nil {
				return out, fmt.Errorf("ERROR: Could not find schema for ID %d in message", msgSchemaId)
			}
			// Appropriately assign the schema
			if deserKey {
				c.KeySchema = schema
			} else {
				c.ValueSchema = schema
			}
		}
		bytes, err := schema.deserializeAvro(avroBytes[5:])
		if err != nil {
			return out, fmt.Errorf("ERROR: Could not deserialize avro: %s", err)
		}
		out = bytes
	}
	return out, nil
}

// retrieveSchemaById reaches out to registryUrl to return the schema given by schemaId.
// It returns the found Schema and any errors if the request encountered.
func retrieveSchemaById(registryUrl string, schemaId uint32) (*Schema, error) {
	regUrl, err := url.Parse(registryUrl)
	if err != nil {
		return nil, err
	}
	regUrl.Path = path.Join(regUrl.Path, fmt.Sprintf("/schemas/ids/%d", schemaId))
	resp, err := http.Get(regUrl.String())
	if err != nil {
		log.Error().Msgf("Request to retrieve schema `%s` failed: %v.", regUrl.Path, err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Error().Msgf("Non-%d response code: %d", http.StatusOK, resp.StatusCode)
		return nil, fmt.Errorf("non-%d response code: %d", http.StatusOK, resp.StatusCode)
	}
	// If Schema Registry returns a 404, don't return an error (expected behavior) but return nil for the schema
	if resp.StatusCode == http.StatusNotFound {
		log.Error().Msgf("Schema Registry returned 404.")
		return nil, nil
	}
	var schemaResponse SchemaRegistryResponse
	err = unmarshall.Into(&schemaResponse, resp.Body)
	if err != nil {
		log.Error().Msgf("Could not unmarshal response.")
		return nil, err
	}
	schemaResponse.Id = schemaId
	return newSchema(schemaResponse)
}

// deserializeAvro deserializes an Avro Message. NOTE: Any magic bytes or schema IDs must be removed before deserialization.
// It returns the deserialized bytes and any error encountered.
func (s *Schema) deserializeAvro(avroMsg []byte) ([]byte, error) {
	native, _, err := s.Codec.NativeFromBinary(avroMsg)
	if err != nil {
		return []byte{}, err
	}
	bytes, err := s.Codec.TextualFromNative(nil, native)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

// newSchema creates a new schema by attempting to parse the Schema Registry Response.
func newSchema(schemaRegResp SchemaRegistryResponse) (*Schema, error) {
	codec, err := goavro.NewCodecForStandardJSONFull(schemaRegResp.Schema)
	if err != nil {
		log.Error().Msgf("Error parsing schema: %v", err)
		return nil, err
	}
	return &Schema{
		Raw: schemaRegResp.Schema,
		// Version: schemaRegResp.Version,
		Id:    schemaRegResp.Id,
		Codec: codec,
	}, nil
}
