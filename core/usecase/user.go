package usecase

import (
	errors "websocketchat/cmd/errors"
	"websocketchat/core/domain"
	"websocketchat/core/entity"
	"websocketchat/core/infra/postgres/repository"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	AddUser(ctx *fiber.Ctx, req entity.User) errors.Error
	GetAllUser(ctx *fiber.Ctx) ([]domain.User, errors.Error)
	GetUserById(ctx *fiber.Ctx, id uint) (domain.User, errors.Error)
	UpdateUser(ctx *fiber.Ctx, req entity.User, id uint) errors.Error
	DeleteUser(ctx *fiber.Ctx, id uint) errors.Error
}

type UseCaseUser struct {
	user repository.UserRepository
}

func NewUseCaseUser(user repository.UserRepository) UserService {
	return &UseCaseUser{
		user: user,
	}
}

func (u *UseCaseUser) AddUser(ctx *fiber.Ctx, req entity.User) errors.Error {
	filter := entity.User{
		LastName:  req.LastName,
		FirstName: req.FirstName,
	}

	user, err := u.user.GetUserByName(ctx, filter)
	if err != nil {
		return errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	if user.ID != 0 {
		return errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.EXISTED_OBJECT_CODE,
			errors.EXISTED_OBJECT_MESSAGE,
		)
	}
	err = u.user.AddUser(ctx, req)
	if err != nil {
		return errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return nil
}
func (u *UseCaseUser) GetAllUser(ctx *fiber.Ctx) ([]domain.User, errors.Error) {

	users, err := u.user.GetAll(ctx)
	if err != nil {
		return users, errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return users, nil
}
func (u *UseCaseUser) GetUserById(ctx *fiber.Ctx, id uint) (domain.User, errors.Error) {

	user, err := u.user.GetUserById(ctx, id)
	if err != nil {
		return user, errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return user, nil
}
func (u *UseCaseUser) UpdateUser(ctx *fiber.Ctx, req entity.User, id uint) errors.Error {

	err := u.user.Update(ctx, req, id)
	if err != nil {
		return errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return nil
}
func (u *UseCaseUser) DeleteUser(ctx *fiber.Ctx, id uint) errors.Error {
	err := u.user.Delete(ctx, id)
	if err != nil {
		return errors.NewCustomHttpError(
			fiber.StatusInternalServerError,
			errors.ERROR_DATABASE_CODE,
			err.Error(),
		)
	}
	return nil
}
