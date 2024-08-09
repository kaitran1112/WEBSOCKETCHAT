package entity

type Message struct {
	SenderID       int    `json:"sender_id" gorm:"not null"`
	ReceiverID     int    `json:"receiver_id" gorm:"not null"`
	MessageContent string `json:"message_content" gorm:"type:text;not null"`
}
