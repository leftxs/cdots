package main

import (
	"github.com/onerciller/cdots/templates"
	"github.com/urfave/cli"
	"os"
)

const (
	version              = "1.0.0"
	editorConfigFileName = "editorconfig"
	travisFileName       = "travis"
	gitignorefileName    = "gitignore"
	createCommand        = "create"
)

func HandleEditorConfig(c *cli.Context) {
	editorConfig := &templates.EditorConfig{Name: editorConfigFileName}
	templ := &Templates{EditorConfig: editorConfig}

	file := File{
		Name:      editorConfigFileName,
		Templates: templ,
	}
	file.Write()
}

func HandleTravis(c *cli.Context) {
	travis := &templates.Travis{Name: travisFileName}
	templ := &Templates{Travis: travis}

	file := File{
		Name:      travisFileName,
		Templates: templ,
		Ext:       ".yml",
	}
	file.Write()
}

func HandleGitignore(c *cli.Context) {
	gitignore := &templates.Gitignore{Name: gitignorefileName}
	templ := &Templates{Gitignore: gitignore}

	file := File{
		Name:      gitignorefileName,
		Templates: templ,
	}
	file.Write()
}

func main() {
	app := cli.NewApp()
	app.Usage = "Create dotfiles with cdots"
	app.Version = version
	app.Commands = []cli.Command{
		{
			Name: createCommand,
			Subcommands: []cli.Command{
				{

					Name:   editorConfigFileName,
					Action: HandleEditorConfig,
				}, {
					Name:   travisFileName,
					Action: HandleTravis,
				},

				{
					Name:   gitignorefileName,
					Action: HandleGitignore,
				},
			},
		},
	}

	_ = app.Run(os.Args)

}
