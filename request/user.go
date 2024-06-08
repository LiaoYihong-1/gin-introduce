package request

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Age      int    `json:"age" validate:"required,gt=0"`
	Password string `json:"password" validate:"required"`
}
