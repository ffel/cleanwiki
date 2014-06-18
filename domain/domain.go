package domain

// pointers or not????
// presently, the data content of a page is not
// that heavy (a string and a slice of bytes, slices
// are light).
type PageRepository interface {
	Store(page *Page)
	FindByTitle(title string) *Page
}

// Use []byte or string for Body is an implemetation detail that
// is determined by choices in infrastructure..
type Page struct {
	Title string
	Body  []byte
}
