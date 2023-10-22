package auth

type Payload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
