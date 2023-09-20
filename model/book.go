package model

type Book struct {
	ISBN   string `json:"isbn" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author"`
}
