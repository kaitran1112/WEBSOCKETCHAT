package usecase

import (
	errors "websocketchat/cmd/errors"
	"websocketchat/core/domain"
	"websocketchat/core/entity"
	"websocketchat/core/infra/postgres/repository"

	"github.com/gofiber/fiber/v2"
)

type MessageService interface {
	AddMessage(ctx *fiber.Ctx, req entity.Message) errors.Error
	GetAllMessage(ctx *fiber.Ctx) ([]domain.Message, errors.Error)
	GetMessageById(ctx *fiber.Ctx, id uint) (domain.Message, errors.Error)
	UpdateMessage(ctx *fiber.Ctx, req entity.Message, id uint) errors.Error
	DeleteMessage(ctx *fiber.Ctx, id uint) errors.Error
}

type UseCaseMessage struct {
	message repository.MessageRepository
}

func NewUseCaseMessage(user repository.UserRepository) UserService {
	return &UseCaseUser{
		user: user,
	}
}

func (u *UseCaseMessage) AddMessage(ctx *fiber.Ctx, req entity.Message) errors.Error {
	err := u.message.AddMessage(ctx, req)
	if err != nil {
		return errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return nil
}
func (u *UseCaseMessage) GetAllMessage(ctx *fiber.Ctx) ([]domain.Message, errors.Error) {

	users, err := u.message.GetAll(ctx)
	if err != nil {
		return users, errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return users, nil
}
func (u *UseCaseMessage) GetMessageById(ctx *fiber.Ctx, id uint) (domain.Message, errors.Error) {

	user, err := u.message.GetMessageById(ctx, id)
	if err != nil {
		return user, errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return user, nil
}
func (u *UseCaseMessage) UpdateMessage(ctx *fiber.Ctx, req entity.Message, id uint) errors.Error {

	err := u.message.Update(ctx, req, id)
	if err != nil {
		return errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return nil
}
func (u *UseCaseMessage) DeleteMessage(ctx *fiber.Ctx, id uint) errors.Error {
	err := u.message.Delete(ctx, id)
	if err != nil {
		return errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return nil
}
