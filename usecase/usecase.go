package usecase

import (
	"fmt"

	"github.com/ffel/cleanwiki/domain"
)

// Wiki is the Interactor equivalent
type Wiki struct {
	logger Logger
	pages  domain.PageRepository
}

// Add implements use case "add a page to the wiki"
func (wiki *Wiki) Add(page *domain.Page) error {
	wiki.logger.Log(fmt.Sprintf("add page %q", page.Title))
	wiki.pages.Store(page)

	return nil
}

type Logger interface {
	Log(message string) error
}
