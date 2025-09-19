package llm

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/models"
)

func BuildPromptFromGraph(nodes []models.Node, edges []models.Edge, targetNodeId uuid.UUID) string {
	edgeGraph := make(map[string][]string)
	for _, e := range edges {
		edgeGraph[e.Source.String()] = append(edgeGraph[e.Source.String()], e.Target.String())
	}

	idToText := make(map[string]string)
	for _, n := range nodes {
		if n.IsSeedNode {
			idToText[n.ID.String()] = fmt.Sprintf("%s (Central Topic)", n.Text)
		} else {
			idToText[n.ID.String()] = n.Text
		}
	}

	var sb strings.Builder
	sb.WriteString("Graph\n")

	for _, n := range nodes {
		children := edgeGraph[n.ID.String()]
		sb.WriteString(fmt.Sprintf("%s -> ", idToText[n.ID.String()]))

		if len(children) > 0 {
			var childNames []string
			for _, c := range children {
				childNames = append(childNames, idToText[c])
			}
			sb.WriteString(strings.Join(childNames, ", "))
		}

		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("Suggest nodes for: %s", idToText[targetNodeId.String()]))
	return sb.String()
}
