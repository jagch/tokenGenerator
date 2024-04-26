package model

type Token struct {
	Token string `json:"token"`
}

type Tokens struct {
	Tokens []Token `json:"tokens"`
}

type ChToken struct {
	Token Token
	Error error
}
