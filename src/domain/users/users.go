package users

//Response will be form in this type of struct
type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

//Request will be form in this type of struct
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
