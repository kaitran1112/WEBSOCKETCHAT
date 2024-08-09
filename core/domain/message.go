package domain

import "time"

type Message struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	SenderID       uint      `json:"sender_id" gorm:"not null"`
	ReceiverID     uint      `json:"receiver_id" gorm:"not null"`
	MessageContent string    `json:"message_content" gorm:"type:text;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}
