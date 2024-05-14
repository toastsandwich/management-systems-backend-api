package models

type Book struct {
	// thumbnail
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	PageCount   int    `json:"pageCount"`
	Cost        int    `json:"cost"`
}
