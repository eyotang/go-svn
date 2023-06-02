package svnlook

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/eyotang/go-svn/moment"
)

var (
	regx    *regexp.Regexp
	actions = []string{"A", "D", "_U", "UU", "U"}
)

type Changed struct {
	Action string `json:"action"`
	Path   string `json:"path"`
}

func (c *Changed) String() string {
	buf, _ := json.Marshal(c)
	return string(buf)
}

func (c *Changed) Parse(line string) (err error) {
	fields := regx.FindStringSubmatch(line)
	if len(fields) < 1 || len(fields[0]) < 3 {
		err = errors.New("格式错误")
		return
	}
	c.Action = fields[1]
	c.Path = fields[2]
	return
}

func (look *Look) Changed(when moment.When, rev string) (changed []*Changed, err error) {
	var (
		ok     bool
		result []Result
		action = moment.ActionFlag(when)
	)

	if result, err = look.RunMarshaled("changed", []string{action, rev}); err != nil {
		return
	}
	if len(result) <= 0 {
		changed = append(changed, &Changed{})
		return
	}

	for idx := range result {
		var change *Changed
		v := result[idx]
		if change, ok = v.(*Changed); !ok {
			err = errors.Errorf("Transfer type failed! %T => %T", v, change)
			return
		}
		changed = append(changed, change)
	}

	return
}

func init() {
	regx = regexp.MustCompile(`^([` + strings.Join(actions, "|") + `]+)\s+(.+)$`)
}
