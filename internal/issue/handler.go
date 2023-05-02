package issue

import (
	"fmt"
	"github.com/charmbracelet/glamour"
)

type ErrPrettyPrinter interface {
	PrettyPrint()
}

type GlamourErrPrettyPrinter struct {
	issue *Issue
}

func NewErrPrettyPrinter(issue *Issue) *GlamourErrPrettyPrinter {
	return &GlamourErrPrettyPrinter{issue: issue}
}

func (g *GlamourErrPrettyPrinter) PrettyPrint() {
	out, err := glamour.Render(string(g.issue.MarkdownMsg()), "light")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func Handle(err error, printer ErrPrettyPrinter) {
	if err == nil {
		return
	}

	if printer != nil {
		printer.PrettyPrint()
		panic(nil)
	}

	panic(err)
}
