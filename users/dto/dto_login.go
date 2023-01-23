package dto

type LoginDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	//Id       int    `json:"id"`
}
