package svn

import (
	"path/filepath"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMkdir_Create(t *testing.T) {
	var (
		path = "Game01/Art/Work/abc"
		dir  = strings.ReplaceAll(filepath.Dir(path), "\\", "/")
	)

	Convey("test Mkdir functions", t, func() {
		repo, err := NewRemoteRepo(repoURL, username, password)
		So(err, ShouldBeNil)
		So(repo, ShouldNotBeNil)

		// 查询路径，应该存在
		info, err := repo.Info(path)
		So(err, ShouldBeNil)
		So(info.Path, ShouldEqual, filepath.Base(path))
		So(info.RelativeURL, ShouldEqual, "^/"+path)

		// 删除路径，应该成功
		err = repo.Delete(dir, "")
		So(err, ShouldBeNil)

		// 创建路径
		err = repo.Mkdir(path, "")
		So(err, ShouldBeNil)
	})
}
