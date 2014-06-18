package domain

// pointers or not????
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
