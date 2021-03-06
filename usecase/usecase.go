// Package `usecase` used `domain` with an application in mind.
package usecase

import (
	"fmt"

	"github.com/ffel/cleanwiki/domain"
)

// Notes is the Interactor equivalent
// At this point I am not sure I have to do with a Wiki,
// a Wiki is a possible application for the Notes
//
// Logger and Pages need to be public for these will be
// injected outside the package (main probably)
type Notes struct {
	Logger Logger
	Pages  domain.PageRepository
}

// Add implements use case "add a page to the notes"
//
// I have decided against "domain.Page" as a parameter for
// this forces also higher layers to know about this concept
//
// This decision seems to be quite ok as it prevents
// tests in `interfaces` from depending on `domain.Page`.
// This suggests that `infrastructure` does not need to know
// about `domain.Page`.
func (notes *Notes) Add(title string, content []byte) error {
	page := domain.Page{Title: title, Body: content}
	notes.Logger.Log(fmt.Sprintf("add page %q", page.Title))
	err := notes.Pages.Store(page)

	return err
}

// Read implements use case "read content by title"
//
// in case there is no contents, an empty slice of bytes
// is returned
func (notes *Notes) Read(title string) ([]byte, error) {
	// we must consider the fact that the implementations FindByTitle
	// may return nil or a zero value Page
	// see e.g. http://tour.golang.org/#42
	//
	// I think it is better that the FindByTitle returns an
	// error instead of an (partially) uninitialized value
	page, err := notes.Pages.FindByTitle(title)

	if err != nil {
		println("problematic page object")
		return make([]byte, 0), err
	}

	return page.Body, nil
}

type Logger interface {
	Log(message string) error
}
