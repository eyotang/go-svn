package svn

func (repo *Repo) Delete(target string, comment string) (err error) {
	var (
		path string
	)
	if path, err = repo.joinPath(target); err != nil {
		return
	}
	if len(comment) <= 0 {
		comment = "Deleting by command."
	}
	if _, err = repo.Output([]string{"delete", "-m", comment, "--force", path}); err != nil {
		return
	}
	return
}
