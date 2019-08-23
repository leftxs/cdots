package templates

import (
	"bytes"
	"text/template"
)

func TemplateToString(templateName string, str string, v interface{}) string {
	var buffer bytes.Buffer

	buffer.Reset()

	tpl := template.New(templateName)

	tpl = template.Must(tpl.Parse(str))
	_ = tpl.Execute(&buffer, v)
	return buffer.String()
}
