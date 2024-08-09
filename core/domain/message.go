package domain

type Message struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	SenderID       int    `json:"sender_id" gorm:"not null"`
	ReceiverID     int    `json:"receiver_id" gorm:"not null"`
	MessageContent string `json:"message_content" gorm:"type:text;not null"`
	CreatedAt      int64  `json:"created_at" gorm:"autoCreateTime"`
}
