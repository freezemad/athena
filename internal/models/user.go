package models

type User struct {
  modelImpl
  UserName     string
  Email        string
  HashPassword string
}

func NewUser(userName string, email string, hashPassword string) *User {
  u := &User{
    UserName:     userName,
    Email:        email,
    HashPassword: hashPassword,
  }

  return u
}
