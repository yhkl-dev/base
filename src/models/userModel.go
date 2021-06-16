package models

import "fmt"

type UserModel struct {
	UserID   int `uri:"id" binding:"required,gt=0"`
	UserName string
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (s *UserModel) String() string {
	return fmt.Sprintf("UserID: %d, UserName: %s", s.UserID, s.UserName)
}
