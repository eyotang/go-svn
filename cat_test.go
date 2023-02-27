package svn

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCat_Content(t *testing.T) {
	var (
		path = "Game01/Art/b.txt"
	)

	Convey("test Cat functions", t, func() {
		repo, err := NewRemoteRepo(repoURL, username, password)
		So(err, ShouldBeNil)
		So(repo, ShouldNotBeNil)

		out, err := repo.Cat(path, "")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeEmpty)

		out, err = repo.Cat(path, "6")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeEmpty)
	})
}
