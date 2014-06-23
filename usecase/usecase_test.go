package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ffel/cleanwiki/domain"
)

// it should be possible to test usecases and domain independently
// however, we need some mockups (or [mock objects](http://en.wikipedia.org/wiki/Mock_object))
// for the missing implementations of the interfaces in Notes

func ExampleAdd() {
	notes := Notes{&logger{}, NewRepo()}

	notes.Add("title", []byte("hi world"))

	// output:
	//
	// add page "title"
}

func ExampleFindByTitle() {
	notes := Notes{&logger{}, NewRepo()}

	notes.Add("title", []byte("hi planet earth"))

	body, _ := notes.Read("title")
	fmt.Printf("%q\n", body)

	body, _ = notes.Read("noSuchTitle")
	fmt.Printf("%q\n", body)

	// output:
	//
	// add page "title"
	// "hi planet earth"
	// ""
}

func TestRead(t *testing.T) {
	notes := Notes{&logger{}, NewRepo()}

	notes.Add("title", []byte("hi planet earth"))

	_, err := notes.Read("title")

	if err != nil {
		t.Error("expected no error")
	}

	_, err = notes.Read("noSuchTitle")

	if err == nil {
		t.Error("expected an error")
	}
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
	pages map[string]domain.Page
}

func NewRepo() *repo {
	r := repo{}
	r.pages = make(map[string]domain.Page)

	return &r
}

// see http://talks.golang.org/2012/10things.slide#8 and
// http://stackoverflow.com/questions/16742331/how-to-mock-abstract-filesystem-in-go
// however, a map is enough here.

func (r *repo) Store(page domain.Page) error {
	r.pages[page.Title] = page

	return nil
}

func (r *repo) FindByTitle(title string) (domain.Page, error) {
	// in case title does not exist, a zero Page object is returned,
	// see http://tour.golang.org/#42

	page, ok := r.pages[title]

	if ok {
		return page, nil
	}

	return domain.Page{}, errors.New("no page found")
}
