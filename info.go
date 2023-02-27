package svn

import (
	"encoding/xml"
	"github.com/pkg/errors"
)

type Info struct {
	XMLName xml.Name   `xml:"info" json:"-"`
	Entry   *InfoEntry `json:"entry"`
}

type InfoEntry struct {
	XMLName xml.Name `xml:"entry" json:"-"`
	Kind    string   `xml:"kind,attr" json:"kind"`
	Path    string   `xml:"path,attr" json:"path"`
	URL     string   `xml:"url" json:"url"`

	RelativeURL string `xml:"relative-url" json:"relativeURL"`
	Commit      Commit `json:"commit"`
}

func (repo *Repo) Info(target string) (entry *InfoEntry, err error) {
	var (
		info   *Info
		result any
		path   string
	)
	if path, err = repo.joinPath(target); err != nil {
		return
	}
	if result, err = repo.RunMarshaled("info", []string{path}); err != nil {
		return
	}
	if info, _ = result.(*Info); info == nil {
		err = errors.Errorf("类型转换错误, 期望: %T, 实际: %T", info, result)
		return
	}
	entry = info.Entry
	return
}
