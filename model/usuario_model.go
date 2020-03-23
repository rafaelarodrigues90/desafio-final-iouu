package model

// Usuario ...
type Usuario struct {
	ID        string `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
}

// Usuarios ...
type Usuarios []Usuario
