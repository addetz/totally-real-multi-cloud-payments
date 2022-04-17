package handlers

type Request struct {
	UserID   string `json:"user_id"`
	ClientID string `json:"client_id"`
	From     string `json:"from_acct"`
	To       string `json:"to_acct"`
	Amount   int64  `json:"amount"`
}
