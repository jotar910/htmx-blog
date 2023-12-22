package components

import (
	"bytes"
	"context"
	std_html "html"
	"io"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yosssi/gohtml"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"

	"github.com/jotar910/buzzer-cms/pkg/logger"
)

func readArticle(filename string) string {
	path := os.Getenv("ARTICLES_FOLDER") + filename + "/main.md"
	f, err := os.ReadFile(path)
	if err != nil {
		logger.L.Errorf("reading article file: %s", path)
		return ""
	}
	return string(f)
}

func rawHTML(html string) templ.Component {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
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

func escapeHTML(source string) string {
	source = gohtml.Format(source)
	source = std_html.EscapeString(source)
	breakline := "<br/>"
	source = strings.Join(
		mapArray(
			strings.Split(source, "\n"),
			func(line string, _ int) string { return strings.ReplaceAll(line, " ", "&nbsp;") + breakline },
		),
		"",
	)
	return source
}

func mapArray[T any, R any](arr []T, cb func(T, int) R) []R {
	res := make([]R, len(arr))
	for i, v := range arr {
		res[i] = cb(v, i)
	}
	return res
}
