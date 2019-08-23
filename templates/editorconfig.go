package templates

type EditorConfig struct {
	Name string
}

func (e *EditorConfig) Template() string {
	return `root = true
[*]
indent_style = tab
end_of_line = lf
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true`

}

func (e *EditorConfig) Render() string {
	return TemplateToString(e.Name, e.Template(), e)
}
