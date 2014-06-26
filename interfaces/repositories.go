// Package `interfaces` implements domain details with external
// solutions in mind e.g. way to store data.
package interfaces

import (
	"io/ioutil"

	"github.com/ffel/cleanwiki/domain"
)

// We define an interface here which for now extends `ReadWriter`
// with nothing.
//
//
type StorageHandler interface {
	ReadWriter
}

// PageFileRepo is a repo that stores notes in files.
//
// In fact, the name is not good enough, `ReadWriterRepo` would
// be more appropriate.
//
// In package `infrastruct` we instantiate a PageFileRepo which
// will provide the ReadWriter.  As such, that package is free
// to use any combination of reader writer chaining.
type PageFileRepo struct {
	handler StorageHandler
}

// Store implements domain.PageRepository.Store
// as a one file per page repository
//
// In usecase.Add I tried to hide domain.Page, here is it again...
// You have to keep in mind that this method is here because of
// `usecase`, not because of `infra`!
//
// In other words, defined in `domain`, used in `usecases` and
// implemented in `interfaces`
func (repo *PageFileRepo) Store(page *domain.Page) error {
	filename := page.Title + ".txt"

	return ioutil.WriteFile(filename, page.Body, 0600)
}

// FindByTitle implements domain.PageRepository.FindByTitle
// as a one file per page repository
//
// In usecase.Read I tried to hide domain.Page, here it is again ...
//
// see Store.
func (repo *PageFileRepo) FindByTitle(title string) (*domain.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	return &domain.Page{Title: title, Body: body}, err
}
