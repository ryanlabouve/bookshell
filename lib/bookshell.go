package bookshell

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
)

func Load() error {
	DIRECTORY_NAME := "security-grimoire"

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fullDirectoryPath := path.Join(currentDirectory, DIRECTORY_NAME)

	// Create folder if doesn't exist
	info, err := os.Stat(DIRECTORY_NAME)
	log.Println(info)
	if os.IsNotExist(err) {
		log.Printf("Folder not found, creating folder for repo")
		repoDirError := os.MkdirAll(DIRECTORY_NAME, 0755)
		if repoDirError != nil {
			log.Fatal(repoDirError)
		}
	}

	// Clone the repo if it doesn't exist
	info, err = os.Stat(path.Join(DIRECTORY_NAME, ".git"))
	log.Println(info)
	if os.IsNotExist(err) {
		log.Printf("Repo not found, cloning")
		repo, err := git.PlainClone(fullDirectoryPath, false, &git.CloneOptions{
			URL:      "https://github.com/ryanlabouve/security-grimoire",
			Progress: os.Stdout,
		})

		fmt.Println(repo)
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Printf("Repo cloned")
	}

	// Pull the repo
	repo, err := git.PlainOpen(fullDirectoryPath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Pulling repo")
	return nil
}
