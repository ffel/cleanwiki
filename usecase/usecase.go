package usecase

import (
	"fmt"

	"github.com/ffel/cleanwiki/domain"
)

// Notes is the Interactor equivalent
// At this point I am not sure I have to do with a Wiki,
// a Wiki is a possible application for the Notes
type Notes struct {
	logger Logger
	pages  domain.PageRepository
}

// Add implements use case "add a page to the notes"
//
// I have decided against "domain.Page" as a parameter for
// this forces also higher layers to know about this concept
func (notes *Notes) Add(title string, content []byte) error {
	page := domain.Page{Title: title, Body: content}
	notes.logger.Log(fmt.Sprintf("add page %q", page.Title))
	notes.pages.Store(&page)

	return nil
}

type Logger interface {
	Log(message string) error
}
