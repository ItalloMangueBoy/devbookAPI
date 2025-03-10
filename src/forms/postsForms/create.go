package postsforms

import (
	"devbookAPI/src/auth"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Create struct {
	Content  string `json:"content,omitempty"`
	AuthorID int64  `json:"author_id,omitempty"`
}

var err error

// Prepare: prepare post data to be created
func (form *Create) Prepare(r *http.Request) error {
	// decode request body
	if err = json.NewDecoder(r.Body).Decode(&form); err != nil {
		return err
	}

	// get authenticated user ID
	if form.AuthorID, err = auth.GetAuthenticatedId(r); err != nil {
		return err
	}

	// format and validate form
	return form.format().validate()
}

func (form *Create) format() *Create {
	// content formatting
	form.Content = strings.TrimSpace(form.Content)

	return form
}

func (form *Create) validate() error {
	// content validation
	if form.Content == "" {
		return errors.New("insert post content")
	}
	if len(form.Content) > 200 {
		return errors.New("post content too long")
	}

	// authorID validation
	if form.AuthorID == 0 {
		return errors.New("author not identified")
	}

	// return success
	return nil
}
