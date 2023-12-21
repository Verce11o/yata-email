package domain

type IncomingMailRequest struct {
	Type string `json:"type"`
	To   string `json:"to"`
	Code string `json:"code"`
}
