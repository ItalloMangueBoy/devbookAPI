package forms

import (
	userhelper "devbookAPI/src/helper/userHelper"
	"encoding/json"
	"fmt"
	"net/http"
)

// PasswordForm: struct for password form
type PasswordForm struct {
	Password        string `json:"password,omitempty"`
	NewPassword     string `json:"new_password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

// Prepare: prepare form data for operations
func (form *PasswordForm) Prepare(r *http.Request, authID int64) error {
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		return err
	}

	return form.format().validate(authID)
}

func (form *PasswordForm) format() *PasswordForm {
	// for future formatting
	return form
}

func (form *PasswordForm) validate(authID int64) error {
	// check actual password
	if form.Password == "" {
		return fmt.Errorf("insert your actual password")
	}

	if err := userhelper.CheckUserPassword(authID, form.Password); err != nil {
		return err
	}

	// check if new password is valid
	if form.NewPassword == "" {
		return fmt.Errorf("insert your new password")
	}

	if len(form.NewPassword) > 50 {
		return fmt.Errorf("new password is too long")
	}

	// check password confirmation
	if form.NewPassword != form.ConfirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	// returns
	return nil
}
