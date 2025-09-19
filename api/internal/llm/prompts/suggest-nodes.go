package prompts

import (
	_ "embed"

	"github.com/neilchetna/mind-mapper/pkg"
)

//go:embed suggest-nodes.tmpl
var suggestNodesTemplate string

type SuggestNodesData struct {
	Graph string
}

func SuggestNodesPrompt(graph string) (string, error) {
	suggestNodesData := SuggestNodesData{graph}
	return pkg.RenderTemplate(suggestNodesTemplate, suggestNodesData)
}
