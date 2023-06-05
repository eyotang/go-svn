package svnlook

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type Look struct {
	binary string
	repo   string
}

type Result interface {
	String() string
}

func NewLook(repo string) *Look {
	return &Look{binary: "/usr/bin/svnlook", repo: repo}
}

func (look *Look) Output(command string, args []string) (out []byte, err error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	b := look.binary
	if !strings.Contains(b, "/") {
		b, _ = exec.LookPath(b)
	}
	cmd := exec.Cmd{
		Path:   b,
		Args:   []string{look.binary, command, look.repo},
		Stdout: &stdout,
		Stderr: &stderr,
		Env:    []string{"LANG=zh_CN.UTF-8"},
	}
	cmd.Args = append(cmd.Args, args...)

	if err = cmd.Run(); err != nil {
		err = errors.Wrap(err, stderr.String())
		return
	}
	out = stdout.Bytes()
	return
}

func (look *Look) RunMarshaled(command string, args []string) (result []Result, err error) {
	var out []byte
	if out, err = look.Output(command, args); err != nil {
		return
	}
	buf := bytes.NewBuffer(out)
	if result, err = Decode(command, buf); err != nil {
		return
	}
	return
}

func Decode(command string, buf *bytes.Buffer) (result []Result, err error) {
	content := strings.Trim(buf.String(), "\n")
	scanner := bufio.NewScanner(buf)
	switch command {
	case "log":
		l := &Log{Content: content}
		result = append(result, l)
	case "author":
		author := &Author{Name: strings.TrimSpace(content)}
		result = append(result, author)
	case "changed":
		for scanner.Scan() {
			line := scanner.Text()
			changed := &Changed{}
			if err = changed.Parse(line); err != nil {
				continue
			}
			result = append(result, changed)
		}
	case "dirs-changed":
		for scanner.Scan() {
			line := scanner.Text()
			dir := &DirsChanged{Path: line}
			result = append(result, dir)
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
