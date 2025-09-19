package dto

type SuggestNodesPayload struct {
	ChartId string `json:"chartId"`
	NodeId  string `json:"nodeId"`
	UserId  string `json:"userId"`
}
