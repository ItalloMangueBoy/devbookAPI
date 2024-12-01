package model

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User: Represents an API user
type User struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// User.Format: format user data before database operations
func (user *User) Format() *User {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	return user
}

// User.Validate: valid user data before database operations
func (user *User) Validate(operation string) error {
	// Username validation
	if user.Name == "" {
		return errors.New("insert your username")
	}
	if len(user.Name) > 50 {
		return errors.New("username too long")
	}

	// Nickname validation
	if user.Nick == "" {
		return errors.New("insert your nickname")
	}
	if len(user.Nick) > 50 {
		return errors.New("nickname too long")
	}

	// Email validation
	if user.Email == "" {
		return errors.New("insert your email")
	}
	if len(user.Email) > 50 {
		return errors.New("email too long")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return err
	}
	if err := checkmail.ValidateHost(user.Email); err != nil {
		return err
	}

	// Password validation
	if operation == "register" {
		if user.Password == "" {
			return errors.New("insert your password")
		}
		if len(user.Password) > 50 {
			return errors.New("password too long")
		}
	}

	// Returns
	return nil
}

// User.Prepare: format and valid user data before database operations
func (user *User) Prepare(r *http.Request, operation string) error {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	return user.Format().Validate(operation)
}
