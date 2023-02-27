package svn

func (repo *Repo) Mkdir(target string, comment string) (err error) {
	var (
		path string
	)
	if path, err = repo.joinPath(target); err != nil {
		return
	}
	if len(comment) <= 0 {
		comment = "Create by command."
	}
	if _, err = repo.Output([]string{"mkdir", "-m", comment, "--parents", path}); err != nil {
		return
	}
	return
}
