package model

type User struct {
	ID    int    `gorM:"int64" json:"id"`
	Name  string `gorM:"varchar(255)" json:"name"`
	Email string `gorM:"varchar(255)" json:"email"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}
