package svnlook

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/eyotang/go-svn/moment"
)

type DirsChanged struct {
	Path string `json:"path"`
}

func (c *DirsChanged) String() string {
	buf, _ := json.Marshal(c)
	return string(buf)
}

func (look *Look) DirsChanged(when moment.When, rev string) (dirs []*DirsChanged, err error) {
	var (
		ok     bool
		result []Result
		action = moment.ActionFlag(when)
	)

	if result, err = look.RunMarshaled("dirs-changed", []string{action, rev}); err != nil {
		return
	}
	if len(result) <= 0 {
		dirs = append(dirs, &DirsChanged{})
		return
	}

	for idx := range result {
		var dir *DirsChanged
		v := result[idx]
		if dir, ok = v.(*DirsChanged); !ok {
			err = errors.Errorf("Transfer type failed! %T => %T", v, dir)
			return
		}
		dirs = append(dirs, dir)
	}

	return
}
