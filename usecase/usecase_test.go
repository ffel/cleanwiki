package usecase

import (
	"fmt"

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

	fmt.Printf("%q\n", notes.Read("title"))
	fmt.Printf("%q\n", notes.Read("noSuchTitle"))

	// output:
	//
	// add page "title"
	// "hi planet earth"
	// ""
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
	pages map[string]*domain.Page
}

func NewRepo() *repo {
	r := repo{}
	r.pages = make(map[string]*domain.Page)

	return &r
}

// see http://talks.golang.org/2012/10things.slide#8 and
// http://stackoverflow.com/questions/16742331/how-to-mock-abstract-filesystem-in-go
// however, a map is enough here.

func (r *repo) Store(page *domain.Page) {
	println("store", page.Title)
	r.pages[page.Title] = page
}

func (r *repo) FindByTitle(title string) *domain.Page {
	println("retrieve", title)
	// in case title does not exist, a zero Page object is returned,
	// see http://tour.golang.org/#42
	return r.pages[title]
}
