package models

type Book struct {
	// thumbnail
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	PageCount   int    `json:"pageCount"`
	Cost        int    `json:"cost"`
}

