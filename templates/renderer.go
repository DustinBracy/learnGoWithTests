package renderer

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/yuin/goldmark"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func (p Post) SanitizedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

type PostRenderer struct {
	templ *template.Template
}

type PostViewModel struct {
	Title, SanitizedTitle, Description, Body string
	Tags                                     []string
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (pr *PostRenderer) Render(w io.Writer, post Post) error {
	return pr.templ.ExecuteTemplate(w, "blog.gohtml", post)

}

func (pr *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return pr.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, post Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if post.Body != "" {
		html, err := markdownToHTML(post.Body)
		if err != nil {
			return err
		}
		post.Body = html
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}

func markdownToHTML(markdown string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
