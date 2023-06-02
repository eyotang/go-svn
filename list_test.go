package svn

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestList_Files(t *testing.T) {
	var (
		project = "Game01"
		dirs    = []string{"Art", "Video"}
		actual  []string
	)

	Convey("test List functions", t, func() {
		repo, err := NewRemoteRepo(repoURL, username, password)
		So(err, ShouldBeNil)
		So(repo, ShouldNotBeNil)

		elements, err := repo.List(project)
		So(err, ShouldBeNil)
		So(elements, ShouldNotBeEmpty)

		for _, e := range elements.Entries {
			actual = append(actual, e.Name)
			fmt.Println(e)
		}
		So(actual, ShouldResemble, dirs)
	})
}
