package email

type emailSchema struct {
	To string `json:"to"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Tags []string `json:"tags"`
}