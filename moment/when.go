package moment

import "github.com/pkg/errors"

type When int

const (
	_When When = iota
	PreCommit
	PostCommit
)

func TransferMoment(moment string) (when When, err error) {
	switch moment {
	case "pre", "pre-commit":
		when = PreCommit
	case "post", "post-commit":
		when = PostCommit
	default:
		err = errors.Errorf("Invalid moment: %s", moment)
	}
	return
}

func ActionFlag(when When) string {
	action := "-r"
	if when == PreCommit {
		action = "-t"
	}
	return action
}
