package svnlook

import (
	"encoding/json"

	"github.com/eyotang/go-svn/moment"
	"github.com/pkg/errors"
)

type Author struct {
	Name string `json:"name"`
}

func (author *Author) String() string {
	buf, _ := json.Marshal(author)
	return string(buf)
}

func (look *Look) Author(when moment.When, rev string) (author *Author, err error) {
	var (
		ok     bool
		result []Result
		action = moment.ActionFlag(when)
	)

	if result, err = look.RunMarshaled("author", []string{action, rev}); err != nil {
		return
	}
	if len(result) <= 0 {
		author = &Author{}
		return
	}

	if author, ok = result[0].(*Author); !ok {
		err = errors.Errorf("Transfer type failed! %T => %T", result[0], author)
		return
	}
	return
}
