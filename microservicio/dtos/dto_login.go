package dtos

type LoginDto struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	//Id       int    `json:"id"`
}
