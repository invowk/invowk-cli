package issue

import (
	"github.com/charmbracelet/glamour"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Id int

const (
	FileNotFoundId Id = iota + 1
	TuiServerStartFailedId
)

type MarkdownMsg string

type HttpLink string

type Renderer interface {
	Render(in string, stylePath string) (string, error)
}

type Issue struct {
	id       Id          // ID used to lookup the issue
	mdMsg    MarkdownMsg // Markdown text that will be rendered
	docLinks []HttpLink  // must never be empty, because we need to have docs about all issue types
	extLinks []HttpLink  // external links that might be useful for the user
}

func (i *Issue) Id() Id {
	return i.id
}

func (i *Issue) MarkdownMsg() MarkdownMsg {
	return i.mdMsg
}

func (i *Issue) DocLinks() []HttpLink {
	return slices.Clone(i.docLinks)
}

func (i *Issue) ExtLinks() []HttpLink {
	return slices.Clone(i.extLinks)
}

func (i *Issue) Render(stylePath string) (string, error) {
	extraMd := ""
	if len(i.docLinks) > 0 || len(i.extLinks) > 0 {
		extraMd += "\n\n"
		extraMd += "## See also: "
		for _, link := range i.docLinks {
			extraMd += "- [" + string(link) + "]"
		}
		for _, link := range i.extLinks {
			extraMd += "- [" + string(link) + "]"
		}
	}
	return render(string(i.mdMsg)+extraMd, stylePath)
}

var (
	render = glamour.Render

	fileNotFoundIssue = &Issue{
		id: FileNotFoundId,
		mdMsg: `
# Dang, we have run into an issue!
We have failed to start our super powered TUI Server due to weird conditions.

## Things you can try to fix and retry
- Run this command  
~~~
$ invowk fix
~~~  
    and try again what you doing before.`,
	}
	issues = map[Id]*Issue{fileNotFoundIssue.Id(): fileNotFoundIssue}
)

func Values() []*Issue {
	return maps.Values(issues)
}

func Get(id Id) *Issue {
	return issues[id]
}
