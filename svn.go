package svn

import (
	"bytes"
	"encoding/xml"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// Output runs p4 and captures stdout.
func (repo *Repo) Output(args []string) (out []byte, err error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	b := repo.binary
	if !strings.Contains(b, "/") {
		b, _ = exec.LookPath(b)
	}
	cmd := exec.Cmd{
		Path:   b,
		Args:   []string{repo.binary},
		Stdout: &stdout,
		Stderr: &stderr,
	}

	if repo.username != "" {
		cmd.Args = append(cmd.Args, "--non-interactive", "--username", repo.username, "--password", repo.password)
	}
	cmd.Args = append(cmd.Args, args...)

	if err = cmd.Run(); err != nil {
		err = errors.Wrap(err, stderr.String())
	}
	out = stdout.Bytes()
	return
}

func (repo *Repo) RunMarshaled(command string, args []string) (result any, err error) {
	var out []byte
	if out, err = repo.Output(append([]string{command, "--xml"}, args...)); err != nil {
		return
	}
	switch command {
	case "list":
		result = &Lists{}
	case "info":
		result = &Info{}
	}
	if err = xml.Unmarshal(out, result); err != nil {
		return
	}
	return
}
