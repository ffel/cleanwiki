package usecase

import (
	"fmt"

	"github.com/ffel/cleanwiki/domain"
)

// it should be possible to test usecases and domain independently
// however, we need some stubs for the missing implementations of
// the interfaces in Notes

func ExampleAdd() {
	notes := Notes{&logger{}, &repo{}}

	notes.Add("title", []byte("hi world"))

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

// see http://talks.golang.org/2012/10things.slide#8 and
// http://stackoverflow.com/questions/16742331/how-to-mock-abstract-filesystem-in-go
// however, a map is enough here.

func (r *repo) Store(page *domain.Page) {

}

func (r *repo) FindByTitle(title string) *domain.Page {
	return nil
}
