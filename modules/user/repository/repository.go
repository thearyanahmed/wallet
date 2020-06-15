package repository

import "fmt"

type Repository struct {}

func (r *Repository) Test() {
	fmt.Println("called here")
}