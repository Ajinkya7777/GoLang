package models

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

/* post request body
{
  "name": "Ajinkya",
  "email": "ajinkya@infobeans.com",
  "password": "password"
}
*/