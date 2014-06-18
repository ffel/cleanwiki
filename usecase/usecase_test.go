package usecase

import (
	"fmt"

	"github.com/ffel/cleanwiki/domain"
)

// it should be possible to test usecases and domain independently
// however, we need some stubs for the missing implementations of
// the interfaces in Wiki

func ExampleAdd() {
	wiki := Wiki{&logger{}, &repo{}}

	page := domain.Page{Title: "title", Body: []byte("hi world")}

	wiki.Add(&page)

	// output:
	//
	// add page "title"
}

// implement Logger

type logger struct {
	log string
}

func (l *logger) Log(message string) error {
	l.log += message + "\n"

	fmt.Println(message)

	return nil
}

// implement domain.PageRepository

type repo struct {
}

func (r *repo) Store(page *domain.Page) {

}

func (r *repo) FindByTitle(title string) *domain.Page {
	return nil
}
