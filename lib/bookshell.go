package bookshell

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
)

func RepoDirectory() string {
	return "security-grimoire"
}

func FullRepoDirectory() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path.Join(currentDirectory, RepoDirectory())
}

func Load() error {

	// Create folder if doesn't exist
	info, err := os.Stat(RepoDirectory())
	log.Println(info)
	if os.IsNotExist(err) {
		log.Printf("Folder not found, creating folder for repo")
		repoDirError := os.MkdirAll(RepoDirectory(), 0755)
		if repoDirError != nil {
			log.Fatal(repoDirError)
		}
	}

	// Clone the repo if it doesn't exist
	info, err = os.Stat(path.Join(RepoDirectory(), ".git"))
	log.Println(info)
	if os.IsNotExist(err) {
		log.Printf("Repo not found, cloning")
		repo, err := git.PlainClone(FullRepoDirectory(), false, &git.CloneOptions{
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
	repo, err := git.PlainOpen(FullRepoDirectory())
	if err != nil {
		log.Println(err)
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		log.Println(err)
		return err
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Pulling repo")
	return nil
}
