package svn

import (
	"net/url"
	"path"
)

type Repo struct {
	binary    string
	remoteURL string
	local     string
	username  string
	password  string
}

// RepoOptionFunc can be used to customize a new GitLab API client.
type RepoOptionFunc func(*Repo) error

func NewRemoteRepo(remote, username, password string) (r *Repo, err error) {
	return NewRepo(remote, "", WithAuth(username, password))
}

func NewRepo(remote, local string, options ...RepoOptionFunc) (r *Repo, err error) {
	r = &Repo{binary: "svn", remoteURL: remote, local: local}
	for _, fn := range options {
		if err = fn(r); err != nil {
			return
		}
	}
	return
}

func WithRemote(remote string) RepoOptionFunc {
	return func(repo *Repo) error {
		repo.remoteURL = remote
		return nil
	}
}

func WithLocal(local string) RepoOptionFunc {
	return func(repo *Repo) error {
		repo.local = local
		return nil
	}
}

func WithAuth(username, password string) RepoOptionFunc {
	return func(repo *Repo) error {
		repo.username = username
		repo.password = password
		return nil
	}
}

func (repo *Repo) joinPath(location string) (output string, err error) {
	var (
		u *url.URL
	)
	if u, err = url.Parse(repo.remoteURL); err != nil {
		return
	}
	u.Path = path.Join(u.Path, location)
	output = u.String()
	return
}
