package main

import (
	//"encoding/json"
	//"fmt"
	//"strings"
	//"testing"

	"docxplate/libs"
)

type User struct {
	Name      string
	Age       int
	Nicknames []string
	Friends   []*User
}


func main() {

user := User{
	Name:      "Alice",
	Age:       27,
	Nicknames: []string{"amber", "AL", "ice"},
	Friends: []*User{
		&User{Name: "Bob", Age: 28},
		&User{Name: "Cecilia", Age: 29},
		&User{Name: "Den", Age: 30},
	},
}

// Template to document
tdoc, _ := docxplate.OpenTemplate("test-data/user.template.docx")
tdoc.Params(user)
tdoc.ExportDocx("MyDoc.docx")

}
