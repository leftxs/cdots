package templates

type Travis struct {
	Name string
}

func (travis *Travis) Template() string {
	return `language: go

go:
  - 1.12.x
env:
  - GO111MODULE=on`

}

func (e *Travis) Render() string {
	return TemplateToString(e.Name, e.Template(), e)
}
