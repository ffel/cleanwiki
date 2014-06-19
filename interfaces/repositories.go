package interfaces

import (
	"io/ioutil"

	"github.com/ffel/cleanwiki/domain"
)

type PageFileRepo struct {
}

// Store implements domain.PageRepository.Store
// as a one file per page repository
//
// In usecase.Add I tried to hide domain.Page, here is it again...
func (repo *PageFileRepo) Store(page *domain.Page) error {
	filename := page.Title + ".txt"

	return ioutil.WriteFile(filename, page.Body, 0600)
}

// FindByTitle implements domain.PageRepository.FindByTitle
// as a one file per page repository
//
// In usecase.Read I tried to hide domain.Page, here it is again ...
func (repo *PageFileRepo) FindByTitle(title string) (*domain.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	return &domain.Page{Title: title, Body: body}, err
}
