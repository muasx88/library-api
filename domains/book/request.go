package book

type AddBookRequestPayload struct {
	Title           string `json:"title" validate:"required,min=5"`
	Author          string `json:"author" validate:"required"`
	PublicationYear int    `json:"publication_year" validate:"required,numeric"`
	Genre           string `json:"genre" validate:"required"`
	ISBN            string `json:"isbn" validate:"required"`
	Stock           int    `json:"stock" validate:"required,numeric"`
}
