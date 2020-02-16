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
	repo *goGit.Repository
	auth *ssh.PublicKeys
)

func InitRepo(path, repoName, username string, key []byte) error {
	var err error
	repo, err = goGit.PlainInit(path, false)
	if err != nil {
		if err == goGit.ErrRepositoryAlreadyExists { // repo already initiated
			repo, err = goGit.PlainOpen(path)
			if err := generateAuth(key); err != nil {
				log.Error().Err(err).Msgf("Error generating key")
				return err
			}

			_ = pullRepoIfExist()
			return nil
		} else {
			log.Error().Err(err).Msgf("Error initiating repository")
		}
		return err
	}

	// repo need to be initiated
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{fmt.Sprintf("git@github.com:%v/%v.git", username, repoName)},
	})
	if err != nil {
		log.Error().Err(err).Msgf("Error creating remote repository config")
		return err
	}

	if err := generateAuth(key); err != nil {
		log.Error().Err(err).Msgf("Error generating key")
		return err
	}

	_ = pullRepoIfExist()
	return nil
}

func generateAuth(key []byte) error {
	var err error
	auth, err = ssh.NewPublicKeys("git", key, "")
	if err != nil {
		log.Error().Err(err).Msgf("Error generating public key")
		return err
	}
	auth.HostKeyCallback = ssh2.InsecureIgnoreHostKey()
	return nil
}

func pullRepoIfExist() error {
	workTree, err := repo.Worktree()
	if err != nil {
		log.Error().Err(err).Msgf("Error getting WorkTree")
		return err
	}

	err = workTree.Pull(&goGit.PullOptions{Auth: auth})
	if err != nil {
		log.Error().Err(err).Msgf("Error pulling the repository - Maybe it is empty?")
		return err
	}

	return nil
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
		log.Debug().Msg("Git status clean -> nothing to commit")
		return nil
	}

	_, err = workTree.Add(".") // add everything to the staging area
	if err != nil {
		log.Error().Err(err).Msgf("Error adding new files to the staging area")
		return err
	}

	_, err = workTree.Commit(fmt.Sprintf("New content from twitter - %v", time.Now().Format("2006-01-02 15:04:05")), &goGit.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: email,
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Error().Err(err).Msgf("Error committing the staging area to the repository")
		return err
	}

	err = repo.Push(&goGit.PushOptions{Auth: auth})
	if err != nil {
		log.Error().Err(err).Msgf("Error pushing the repository")
		return err
	}

	log.Debug().Msg("Successfully pushed commit(s) to the repository")

	return err
}
