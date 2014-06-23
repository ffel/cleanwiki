package domain

// pointers or not????
// presently, the data content of a page is not
// that heavy (a string and a slice of bytes, slices
// are light).
// Additionally, you probably don't want to allow modifying Page
// the moment you received one by its title.
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
