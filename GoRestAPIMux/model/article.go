package model

import (
	"fmt"
)

type Article struct {
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
	ID     string `json:"Id"`
}

func (p Article) Print() {
	fmt.Println("\n ======= \n")
	fmt.Println("Title", p.Title)
	fmt.Println("Author", p.Author)
	fmt.Println("Link", p.Link)
	fmt.Println("ID", p.ID)
}

func (articlePointer *Article) UpdateArticleDetail(title string, author string, link string, id string) {
	(*articlePointer).Title = title
	(*articlePointer).Author = author
	(*articlePointer).Link = link
	(*articlePointer).ID = id
}

func TestPrint() {
	fmt.Println("This is working")
}
