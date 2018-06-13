package main

/**
 *  Struct that represents a book.
 */
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

/*
	The (b Book) part tells Go that this func is related to the Book type,
	i.e. is a method of a book.
*/
func (b Book) toJSON() string {
	return toJSON(b)
}
