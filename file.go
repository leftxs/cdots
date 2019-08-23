package main

import (
	"github.com/onerciller/cdots/templates"
	"io/ioutil"
)

const (
	permisson    = 0755
	editorconfig = "editorconfig"
	travis       = "travis"
	gitignore    = "gitignore"
)

type Templates struct {
	EditorConfig *templates.EditorConfig
	Travis       *templates.Travis
	Gitignore    *templates.Gitignore
}

type File struct {
	Name      string
	Ext       string
	Templates *Templates
}

func (file *File) getFile() string {
	switch file.Name {
	case editorconfig:
		return file.Templates.EditorConfig.Render()
	case travis:
		return file.Templates.Travis.Render()
	case gitignore:
		return file.Templates.Gitignore.Render()
	default:
		return ""
	}
}

func (file *File) Write() {
	_ = ioutil.WriteFile("."+file.Name+file.Ext, []byte(file.getFile()), permisson)
}
