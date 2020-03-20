package model

type Usuario struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
}

type Usuarios struct {
	Usuarios []Usuario	`json:"usuario"`
}