package models

type User struct {
  modelImpl
  UserName     string  `db:"UserName"`
  Email        string  `db:"Email"`
  HashPassword string  `db:"HashPassword"`
}

func NewUser(userName string, email string, hashPassword string) *User {
  u := &User{
    UserName:     userName,
    Email:        email,
    HashPassword: hashPassword,
  }

  return u
}
