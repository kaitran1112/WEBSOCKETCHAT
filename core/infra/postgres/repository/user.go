package repository

import (
	"websocketchat/core/domain"
	"websocketchat/core/entity"
	"websocketchat/core/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(ctx *fiber.Ctx, user entity.User) error
	GetUserById(ctx *fiber.Ctx, id uint) (domain.User, error)
	GetAll(ctx *fiber.Ctx) (users []domain.User, err error)
	Update(ctx *fiber.Ctx, req entity.User, id uint) error
	Delete(ctx *fiber.Ctx, id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u userRepository) AddUser(ctx *fiber.Ctx, user entity.User) error {
	id := utils.GenerateUniqueKey()
	userSaving := domain.User{
		ID:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
	err := u.DB.Save(userSaving).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) GetUserById(ctx *fiber.Ctx, id uint) (domain.User, error) {
	var user domain.User
	err := u.DB.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (u *userRepository) GetAll(ctx *fiber.Ctx) (users []domain.User, err error) {
	err = u.DB.Find(&users).Error
	return users, err
}

func (u *userRepository) Update(ctx *fiber.Ctx, req entity.User, id uint) error {
	err := u.DB.Model(&domain.User{}).Where("ID = ?", id).Updates(req).Error
	return err
}

func (u *userRepository) Delete(ctx *fiber.Ctx, id uint) error {
	err := u.DB.Delete(&domain.User{}, id).Error
	return err
}
