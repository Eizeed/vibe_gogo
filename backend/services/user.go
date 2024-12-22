package services

import (
	"errors"

	"github.com/Eizeed/vibe_gogo/forms"
	"github.com/Eizeed/vibe_gogo/models"
	"github.com/google/uuid"
)

type UserService struct {};

var userModel = models.UserModel {};

func (s *UserService) GetAll() ([]models.User, error) {
    return userModel.GetAll()
}

func (s *UserService) Create(userData forms.RegisterForm) (models.User, error) {
    _, err := userModel.GetByEmail(userData.Email);
    if err == nil {
        return models.User{}, errors.New("User with this email already exists");
    }
    
    uuid, err := uuid.NewV7();
    if err != nil {
        return models.User{}, errors.New("Failed to create uuid for user");
    }

    user := models.User {
        UUID: uuid,
        Email: userData.Email,
        Username: userData.Username,
        Fullname: userData.Fullname,
        Password: userData.Password,
    };
    res, err := userModel.Create(user);

    return res, nil;
}

func (s *UserService) Login(userData forms.LoginForm) (models.User, string, error) {
    res, err := userModel.GetByEmail(userData.Email);
    if err != nil {
        return models.User{}, "", err;
    }

    if res.Password != userData.Password {
        return models.User{}, "", errors.New("Invalid email or password");
    } else {
        jwt := models.JWT {};
        token, err := jwt.GenToken(res.UUID);
        if err != nil {
            return models.User{}, "", err
        }

        return res, token, nil;
    }
}

func (s *UserService) Update(userData forms.UpdateForm, uuid uuid.UUID) (models.User, error) {
    res, err := userModel.Update(userData, uuid);
    if err != nil {
        return models.User{}, err;
    }
    return res, nil;
}

func (s *UserService) Delete(uuid uuid.UUID) (models.User, error) {
    res, err := userModel.Delete(uuid);
    if err != nil { 
        return models.User{}, err;
    }

    return res, nil;
}














