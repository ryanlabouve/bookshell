package bookshell

import (
	"log"

	"github.com/hashicorp/go-memdb"
)

type Book struct {
	Id     string `json:"id" xml:"id" form:"id" query:"id"`
	Title  string `json:"title" xml:"title" form:"title" query:"title"`
	Author string `json:"author" xml:"author" form:"author" query:"author"`
	Slug   string `json:"slug" xml:"slug" form:"slug" query:"slug"`
}

func InsertBooks(db *memdb.MemDB, books []*Book) {
	txn := db.Txn(true)
	for _, b := range books {
		if err := txn.Insert("book", b); err != nil {
			panic(err)
		}
	}
	txn.Commit()
}

func GetBooks(db *memdb.MemDB) []Book {
	// txn := db.Txn(false) // read-only
	txn := db.Txn(true)
	defer txn.Abort()
	var books []Book

	it, err := txn.Get("book", "id")

	if err != nil {
		// TODO fail better
		log.Println("PANIC19199")
		panic(err)
	}

	for o := it.Next(); o != nil; o = it.Next() {
		log.Println(o)
		book := (o).(Book)
		books = append(books, book)
	}

	log.Println("2")

	return books
}
