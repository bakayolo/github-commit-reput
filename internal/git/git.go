package git

import (
	"fmt"
	"github.com/rs/zerolog/log"
	ssh2 "golang.org/x/crypto/ssh"
	goGit "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"time"
)

var (
	repo      *goGit.Repository
	deployKey []byte
)

func InitRepo(path, repoName, username string, key []byte) error {
	deployKey = key

	var err error
	repo, err = goGit.PlainInit(path, false)
	if err != nil {
		if err == goGit.ErrRepositoryAlreadyExists { // repo already initiated
			repo, err = goGit.PlainOpen(path)
		}
		return err
	}

	// repo need to be initiated
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{fmt.Sprintf("git@github.com:%v/%v.git", username, repoName)},
	})
	return err
}

func CommitAndPushRepo(username, email string) error {
	workTree, err := repo.Worktree()
	if err != nil {
		log.Error().Err(err).Msgf("Error getting WorkTree")
		return err
	}

	status, err := workTree.Status()
	if err != nil {
		log.Error().Err(err).Msgf("Error retrieving status from workTree")
		return err
	}

	if status.IsClean() { // nothing to do
		return nil
	}

	_, err = workTree.Add(".") // add everything to the staging area
	if err != nil {
		log.Error().Err(err).Msgf("Error adding new files to the staging area")
		return err
	}

	_, err = workTree.Commit("New content", &goGit.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: email,
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Error().Err(err).Msgf("Error commiting the staging area to the repository")
		return err
	}

	auth, err := ssh.NewPublicKeys("git", deployKey, "")
	if err != nil {
		log.Error().Err(err).Msgf("Error generating public key")
		return err
	}
	auth.HostKeyCallback = ssh2.InsecureIgnoreHostKey()

	err = repo.Push(&goGit.PushOptions{Auth: auth})
	if err != nil {
		log.Error().Err(err).Msgf("Error pushing the repository")
	}
	return err
}
