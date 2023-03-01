package svn

func (repo *Repo) CatRev(target string, revision string) (out []byte, err error) {
	var (
		path string
	)
	if path, err = repo.joinPath(target); err != nil {
		return
	}
	if len(revision) <= 0 {
		revision = "HEAD"
	}
	if out, err = repo.Output([]string{"cat", "-r", revision, path}); err != nil {
		return
	}
	return
}

func (repo *Repo) Cat(target string) (out []byte, err error) {
	return repo.CatRev(target, "HEAD")
}
