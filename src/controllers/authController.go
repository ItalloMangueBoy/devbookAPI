package controllers

import (
	"database/sql"
	"devbookAPI/src/auth"
	userhelper "devbookAPI/src/helper/userHelper"
	"devbookAPI/src/security"
	"devbookAPI/src/view"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/badoux/checkmail"
)

type loginForm struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (form *loginForm) validate() error {
	// Email validation
	if strings.TrimSpace(form.Email) == "" {
		return errors.New("invalid credentials")
	}
	if err := checkmail.ValidateFormat(form.Email); err != nil {
		return errors.New("invalid credentials")
	}

	// Password validation
	if strings.TrimSpace(form.Password) == "" {
		return errors.New("invalid credentials")
	}
	if len(form.Password) > 50 {
		return errors.New("invalid credentials")
	}

	// Correct
	return nil
}

func (form *loginForm) prepare(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		return err
	}

	return form.validate()
}

// Login: auth an user
func Login(w http.ResponseWriter, r *http.Request) {
	// Recive request data
	var form loginForm

	if err := form.prepare(r); err != nil {
		view.GenErrorTemplate(err).Send(w, 401)
		return
	}

	// Get user who will be logged
	user, err := userhelper.SearchUserByEmail(form.Email)
	if err == sql.ErrNoRows {
		view.GenErrorTemplate(errors.New("invalid credentials")).Send(w, 401)
		return
	}

	if err != nil {
		view.GenErrorTemplate(err).Send(w, 500)
		return
	}

	// Check password
	if !security.ValidPassword(user.Password, form.Password) {
		view.GenErrorTemplate(errors.New("invalid credentials")).Send(w, 401)
	}

	// Gen token
	token, err := auth.GenToken(user)
	if err != nil {
		view.GenErrorTemplate(fmt.Errorf("can not generate user token")).Send(w, 500)
		return
	}

	view.JSON(w, 200, token)
}


