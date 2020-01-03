package models

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint32    `json:"id"`
	Nickname  string    `json:"nickname" validate:"required" `
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Role      string    `json:"role"  validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
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

func (u *LoginUser) Login(db *sqlx.DB) (*User, error) {
	var user = &User{}
	err := db.Get(user, "SELECT * FROM users WHERE email=$1", u.Email)
	if err != nil {
		return &User{}, err
	}
	err = VerifyPass(u.Password, user.Password)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (u *User) SaveUser(db *sqlx.DB) (*User, error) {
	var err error
	hash_pass, err := Hash(u.Password)
	if err != nil {
		return &User{}, err
	}
	u.Password = string(hash_pass)
	user_query := `INSERT INTO users (nickname, email, password, role) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(user_query, u.Nickname, u.Email, u.Password, u.Role)
	if err != nil {
		return &User{}, err
	}
	var user = &User{}
	db.Get(user, "SELECT * FROM users WHERE email=$1", u.Email)
	return user, nil
}

func (u *User) FindAllUsers(db *sqlx.DB) (*[]User, error) {
	var err error
	users := &[]User{}
	err = db.Select(users, "SELECT * FROM users")
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *User) FindUser(db *sqlx.DB, id string) (*User, error) {
	var err error
	user := &User{}
	err = db.Get(user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) DeleteUser(db *sqlx.DB, id string) error {
	var err error
	_, err = db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(db *sqlx.DB, id string) (*User, error) {
	hash_pass, err := Hash(u.Password)
	if err != nil {
		return &User{}, err
	}
	u.Password = string(hash_pass)
	update_query := "UPDATE users SET email=$1,password=$2,nickname=$3,update_at=$4"
	_, err = db.Exec(update_query, u.Password, u.Nickname, u.Email, time.Now())
	if err != nil {
		return &User{}, err
	}
	user := &User{}
	err = db.Get(user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return user, err
	}
	return user, nil
}
