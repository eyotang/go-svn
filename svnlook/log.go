package svnlook

import (
	"encoding/json"
	"github.com/pkg/errors"

	"svn-hook/moment"
)

type Log struct {
	Content string `json:"content"`
}

func (log *Log) String() string {
	buf, _ := json.Marshal(log)
	return string(buf)
}

func (look *Look) Log(when moment.When, rev string) (log *Log, err error) {
	var (
		ok     bool
		result []Result
		action = moment.ActionFlag(when)
	)
	if result, err = look.RunMarshaled("log", []string{action, rev}); err != nil {
		return
	}
	if len(result) <= 0 {
		log = &Log{}
		return
	}

	if log, ok = result[0].(*Log); !ok {
		err = errors.Errorf("Transfer type failed! %T => %T", result[0], log)
		return
	}
	return
}
