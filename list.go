package svn

import (
	"encoding/json"
	"encoding/xml"
	"github.com/pkg/errors"
)

type Lists struct {
	XMLName xml.Name     `xml:"lists" json:"-"`
	Entries []*ListEntry `xml:"list>entry" json:"entries"`
}

type ListEntry struct {
	XMLName xml.Name `xml:"entry" json:"-"`
	Kind    string   `xml:"kind,attr" json:"kind"`
	Name    string   `xml:"name" json:"name"`
	Commit  Commit   `json:"commit"`
}

type Commit struct {
	XMLName xml.Name `xml:"commit" json:"-"`
	Rev     string   `xml:"revision,attr" json:"revision"`
	Author  string   `xml:"author" json:"author"`
	Date    string   `xml:"date" json:"date"`
}

func (lists *Lists) GetEntries() []*ListEntry {
	if lists == nil {
		return nil
	}
	return lists.Entries
}

func (e *ListEntry) String() string {
	buf, _ := json.Marshal(e)
	return string(buf)
}

func (repo *Repo) List(target string) (lists *Lists, err error) {
	var (
		result any
		path   string
	)
	if path, err = repo.joinPath(target); err != nil {
		return
	}
	if result, err = repo.RunMarshaled("list", []string{path}); err != nil {
		return
	}
	if lists, _ = result.(*Lists); lists == nil {
		err = errors.Errorf("类型转换错误, 期望: %T, 实际: %T", lists, result)
		return
	}
	return
}
