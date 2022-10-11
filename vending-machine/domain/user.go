package domain

type User struct {
	ID int `json:"id" pg:"id,pk"`
	Username string `json:"Username" pg:"username"`
	Password string `json:"Password" pg:"password"`
	Deposit int `json:"Deposit" pg:"deposit"`
	Role string `json:"Role" pg:"role"`
}