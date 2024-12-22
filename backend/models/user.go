package models

import (
	"errors"

	"github.com/Eizeed/vibe_gogo/db"
	"github.com/Eizeed/vibe_gogo/forms"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
    UUID        uuid.UUID   `json:"uuid"`
    Email       string      `json:"email"`
    Username    string      `json:"username"`
    Fullname    string      `json:"fullname"`
    Password    string      `json:"password"`
}

type UserModel struct {};

func (m *UserModel) GetAll() ([]User, error) {
    db := db.GetDB();
    
    var users []User;

    db.Find(&users)

    return users, nil;
}

func (m *UserModel) GetById() (string, error) {
    return "returning user by id", nil;
}

func (m *UserModel) GetByEmail(email string) (User, error) {
    db := db.GetDB();

    var user User;

    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return User{}, errors.New("Invalid email or password");
        }
        return User{}, err;
    }

    return user, nil;
}

func (m *UserModel) Create(userData User) (User, error) {
    db := db.GetDB();

    if err := db.Clauses(clause.Returning{}).Create(&userData).Error; err != nil {
        return User{}, err;
    }

    var user User;

    if err := db.Where(userData).First(&user).Error; err != nil {
        return User{}, err;
    }
    

    return user, nil;
}

func (m *UserModel) Update(userData forms.UpdateForm, uuid uuid.UUID) (User, error) {
    db := db.GetDB();

    var user User;

    if err := db.Model(&user).Clauses(clause.Returning{}).Where("uuid = ?", uuid).Updates(userData).Error; err != nil {
        return User{}, err;
    }

    return user, nil;
}

func (m *UserModel) Delete(uuid uuid.UUID) (User, error) {
    db := db.GetDB();

    var user User;

    if res := db.Clauses(clause.Returning{}).Where("uuid = ?", uuid).Delete(&user); res != nil {
        if res.RowsAffected == 0 {
            return User{}, errors.New("User not found");
        }
        if res.Error != nil {
            return User{}, res.Error;
        }
    }

    return user, nil;
}













