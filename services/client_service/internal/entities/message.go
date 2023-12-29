package entities

import (
	"github.com/google/uuid"
	"time"
)

type MessageStatus byte

const (
	StatusPending MessageStatus = iota
	StatusDelivered
)

type Message struct {
	Id             uuid.UUID     `json:"uuid"`
	SentAt         time.Time     `json:"sent_at"`
	Status         MessageStatus `json:"status"`
	ClientId       uuid.UUID     `json:"client_id"`
	CreatedAt      time.Time     `json:"created_at"`
	NotificationId uuid.UUID     `json:"notification_id"`
}

func NewMessage(clientId uuid.UUID, notificationId uuid.UUID) *Message {
	return &Message{
		Id:             uuid.New(),
		SentAt:         time.Now(),
		Status:         StatusPending,
		ClientId:       clientId,
		CreatedAt:      time.Now(),
		NotificationId: notificationId,
	}
}

func (m *Message) Send() {
	m.SentAt = time.Now()
	m.Status = StatusDelivered
}
