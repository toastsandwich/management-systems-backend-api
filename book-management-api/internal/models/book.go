package models

type Book struct {
	// thumbnail
	ID          int    `json:"id"`
	Cost        int    `json:"cost"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	ISBN        string `json:"isbn"`
	PageCount   int    `json:"pageCount"`
	Title       string `json:"title"`
}
