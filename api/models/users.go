package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname" validate:"required" `
	Email     string    `gorm:"size:100;not null;unique" json:"email" validate:"required,email"`
	Password  string    `gorm:"size:100;not null;" json:"password" validate:"required"`
	Role      string    `gorm:"size:100;not null;" json:"role"  validate:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type LoginUser struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPass(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *LoginUser) Login(db *gorm.DB) (*User, error) {
	var user = &User{}
	err := db.Where("email = ?", u.Email).Find(&user).Error
	if err != nil {
		return &User{}, err
	}
	err = VerifyPass(u.Password, user.Password)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	hash_pass, err := Hash(u.Password)
	if err != nil {
		return &User{}, err
	}
	u.Password = string(hash_pass)
	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Find(&users).Error
	if err != nil {
		return &users, err
	}
	return &users, nil
}

func (u *User) FindUser(db *gorm.DB, id string) (*User, error) {
	var err error
	user := &User{}
	err = db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) DeleteUser(db *gorm.DB, id string) error {
	var err error
	err = db.Where("id = ?", id).Delete(User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(db *gorm.DB, id string) (*User, error) {
	hash_pass, err := Hash(u.Password)
	if err != nil {
		return &User{}, err
	}
	u.Password = string(hash_pass)
	err = db.Model(&User{}).Where("id = ?", id).Updates(map[string]interface{}{

		"password":  u.Password,
		"nickname":  u.Nickname,
		"email":     u.Email,
		"update_at": time.Now(),
	}).Error
	if err != nil {
		return &User{}, err
	}
	err = db.Model(&User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
