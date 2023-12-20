// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.432
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"

import "os"
import "io"
import "context"
import "github.com/jotar910/htmx-templ/pkg/logger"
import "github.com/yuin/goldmark"
import "github.com/yuin/goldmark/renderer/html"
import "bytes"
import "github.com/microcosm-cc/bluemonday"

func readArticle(filename string) string {
	f, err := os.ReadFile("public/articles/" + filename + "/main.md")
	if err != nil {
		logger.L.Errorf("reading article file: %s", filename)
		return ""
	}
	return string(f)
}

func rawHTML(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, html)
		return err
	})
}

func mdToHTML(source string) string {
	var buf bytes.Buffer
	err := goldmark.Convert([]byte(source), &buf)
	if err != nil {
		return err.Error()
	}
	html := bluemonday.UGCPolicy().SanitizeBytes(buf.Bytes())
	return string(html)
}

func mdToUnsafeHTML(source string) string {
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	err := md.Convert([]byte(source), &buf)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
