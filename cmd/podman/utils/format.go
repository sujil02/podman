package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"text/template"
)

var basicFunctions = template.FuncMap{
	"json": func(v interface{}) string {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		enc.Encode(v)
		return strings.TrimSpace(buf.String())
	},
}

// GenerateTemplate return a Template having mapped basic functions
func GenerateTemplate(tag, format string) (*template.Template, error) {
	return template.New(tag).Funcs(basicFunctions).Parse(format)
}

// ExecuteTemplateAndFlush basically writes the appropriate formated response based on the template
// and flushes out the output.
func ExecuteTemplateAndFlush(tmpl *template.Template, data interface{}, w io.Writer) error {
	if err := tmpl.Execute(w, data); err != nil {
		return err
	}
	if flusher, ok := w.(interface{ Flush() error }); ok {
		return flusher.Flush()
	}
	return nil
}
