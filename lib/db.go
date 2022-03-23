package bookshell

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gernest/front"
	memdb "github.com/hashicorp/go-memdb"
)

func InitializeDb() *memdb.MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"book": {
				Name: "book",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Slug"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return db
}

func SeedDb(db *memdb.MemDB) {
	// Reset the db on each run

	txn := db.Txn(true)

	filepath.Walk(FullRepoDirectory(), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		r, _ := regexp.Compile(`.+books\/__mdx.+(md|mdx)$`)
		if !r.MatchString(path) {
			return nil
		}

		// Parse each of those files
		m := front.NewMatter()
		m.Handle("---", front.YAMLHandler)
		txt, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		f, body, err := m.Parse(strings.NewReader(string(txt)))
		if err != nil {
			panic(err)
		}

		fmt.Printf("The front matter is:\n%#v\n", f)
		fmt.Printf("The body is:\n%q\n", body)

		book_actual := f["meta"].(map[interface{}]interface{})
		if err := txn.Insert("book", Book{
			Slug:   book_actual["slug"].(string),
			Author: book_actual["author"].(string),
			Title:  book_actual["title"].(string),
		}); err != nil {
			panic(err)
		}
		return nil
	})

	// Put that info in the db
	txn.Commit()
}
