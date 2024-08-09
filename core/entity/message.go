package entity

type Message struct {
	SenderID       uint   `json:"sender_id" gorm:"not null"`
	ReceiverID     uint   `json:"receiver_id" gorm:"not null"`
	MessageContent string `json:"message_content" gorm:"type:text;not null"`
}
