package svnlook

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChanged_Parse(t *testing.T) {
	Convey("test changed Parse", t, func() {

		Convey("List changed", func() {
			line := "A   trunk/vendors/deli/"
			changed := Changed{}
			err := changed.Parse(line)
			So(err, ShouldBeNil)
		})

		Convey("Parse changed", func() {
			lines := []string{
				"A trunk/vendors/deli/pickle.txt",
				"U trunk/vendors/baker/bagel.txt",
				"_U trunk/vendors/baker/croissant.txt",
				"UU trunk/vendors/baker/pretzel.txt",
				"D trunk/vendors/baker/baguette.txt",
			}
			changed := Changed{}
			for _, line := range lines {
				err := changed.Parse(line)
				So(err, ShouldBeNil)
			}
		})
	})
}
