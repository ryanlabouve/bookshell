package bookshell

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func Load() error {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	repo, error := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/ryanlabouve/security-grimoire",
		Progress: os.Stdout,
	})

	if error != nil {
		return error
	}

	fmt.Println(repo)
	return nil
}
