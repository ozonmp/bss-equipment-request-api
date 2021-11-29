package sender

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/model"
	"google.golang.org/protobuf/proto"
	"strconv"
)

// TopicType is a type of topic
type TopicType string

func (t TopicType) String() string {
	return string(t)
}

const (
	// CreatedTopic is a "created item" type of events
	CreatedTopic TopicType = "EQUIPMENT_REQUEST_CREATED"
	// UpdatedEquipmentIDTopic is a "updated equipment id of item" type of events
	UpdatedEquipmentIDTopic = "EQUIPMENT_REQUEST_UPDATED_EQUIPMENT_ID"
	// UpdatedStatusTopic is a "updated status of item" type of events
	UpdatedStatusTopic = "EQUIPMENT_REQUEST_UPDATED_STATUS"
	// RemovedTopic is a "removed item" type of events
	RemovedTopic = "EQUIPMENT_REQUEST_DELETED"
	// DefaultTopic is a "removed item" type of events
	DefaultTopic = "EQUIPMENT_REQUEST_DEFAULT"
)

// EventSender is sender of events
type EventSender interface {
	Send(equipmentRequest *model.EquipmentRequestEvent) error
}

type eventSender struct {
	ctx      context.Context
	brokers  []string
	producer sarama.SyncProducer
}

// NewEventSender returns EventSender interface
func NewEventSender(ctx context.Context, brokers []string) (EventSender, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &eventSender{
		ctx:      ctx,
		brokers:  brokers,
		producer: producer,
	}, nil
}

func (e eventSender) Send(equipmentRequest *model.EquipmentRequestEvent) error {
	topic, ok := parseTopic(equipmentRequest)
	if !ok {
		return errors.New("parseTopic: unable to parse topic from event type")
	}

	pbMessage, err := model.ConvertEquipmentRequestEventToPb(equipmentRequest)

	if err != nil {
		return err
	}

	message, err := proto.Marshal(pbMessage)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic.String(),
		Key:   sarama.StringEncoder(strconv.FormatUint(equipmentRequest.EquipmentRequestID, 10)),
		Value: sarama.ByteEncoder(message),
	}
	_, _, err = e.producer.SendMessage(msg)

	return err
}

func parseTopic(equipmentRequest *model.EquipmentRequestEvent) (TopicType, bool) {
	switch equipmentRequest.Type {
	case model.Created:
		return CreatedTopic, true
	case model.UpdatedEquipmentID:
		return UpdatedEquipmentIDTopic, true
	case model.UpdatedStatus:
		return UpdatedStatusTopic, true
	case model.Removed:
		return RemovedTopic, true
	default:
		return DefaultTopic, false
	}
}
