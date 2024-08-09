package repository

import (
	"websocketchat/core/domain"
	"websocketchat/core/entity"
	"websocketchat/core/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MessageRepository interface {
	AddMessage(ctx *fiber.Ctx, user entity.Message) error
	GetMessageById(ctx *fiber.Ctx, id uint) (domain.Message, error)
	GetAll(ctx *fiber.Ctx) (messages []domain.Message, err error)
	Update(ctx *fiber.Ctx, req entity.Message, id uint) error
	Delete(ctx *fiber.Ctx, id uint) error
}

type messageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(DB *gorm.DB) MessageRepository {
	return &messageRepository{
		DB: DB,
	}
}

func (u *messageRepository) AddMessage(ctx *fiber.Ctx, message entity.Message) error {
	id := utils.GenerateUniqueKey()
	messageSaving := domain.Message{
		ID:             id,
		SenderID:       message.SenderID,
		ReceiverID:     message.ReceiverID,
		MessageContent: message.MessageContent,
	}
	err := u.DB.Save(messageSaving).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *messageRepository) GetMessageById(ctx *fiber.Ctx, id uint) (domain.Message, error) {
	var message domain.Message
	err := u.DB.Find(&message).Error
	if err != nil {
		return message, err
	}
	return message, nil
}
func (u *messageRepository) GetAll(ctx *fiber.Ctx) (messages []domain.Message, err error) {
	err = u.DB.Find(&messages).Error
	return messages, err
}

func (u *messageRepository) Update(ctx *fiber.Ctx, req entity.Message, id uint) error {
	err := u.DB.Model(&domain.Message{}).Where("ID = ?", id).Updates(req).Error
	return err
}

func (u *messageRepository) Delete(ctx *fiber.Ctx, id uint) error {
	err := u.DB.Delete(&domain.Message{}, id).Error
	return err
}
