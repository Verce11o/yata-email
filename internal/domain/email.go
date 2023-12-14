package domain

type IncomingMailRequest struct {
	To   string `json:"to"`
	Code string `json:"code"`
}
