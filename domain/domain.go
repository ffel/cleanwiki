package domain

// pointers or not????
// presently, the data content of a page is not
// that heavy (a string and a slice of bytes, slices
// are light).
// Additionally, you probably don't want to allow modifying Page
// the moment you received one by its title.
//
// How convenient is it to implement Writer and Reader?
// this might improve the simplicity in `interfaces`
// but what I don't like is that I'll need an artificial
// split between Title and Body, if that's even possible at all ...
//
// The question is whether or not at *this* level you want
// `Title` to identify the body of a note.  One might even
// wonder whether we need a domain at all as soon as
// we use standard interfaces.
//
// The alternative is to use standard interfaces in `usecases`,
// which serves `interfaces`.
type PageRepository interface {
	Store(page Page) error
	FindByTitle(title string) (Page, error)
}

// Use []byte or string for Body is an implemetation detail that
// is determined by choices in infrastructure..
type Page struct {
	Title string
	Body  []byte
}
