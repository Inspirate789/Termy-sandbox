package benchmark

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"gorm.io/gorm"
	"time"
)

type UsersGormRepository struct {
	gormDB *gorm.DB
}

func (r *UsersGormRepository) AddUser(user interfaces.UserDTO) error {
	model := entities.User{
		Name:     user.Name,
		Password: user.Password,
		Role:     user.Role,
	}
	model.SetRegDate(time.Now())

	return r.gormDB.Create(&model).Error
}

func (r *UsersGormRepository) GetUser(request entities.AuthRequest) (*entities.User, error) {
	var user entities.User
	err := r.gormDB.Where(&entities.User{Name: request.Username, Password: request.Password}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
