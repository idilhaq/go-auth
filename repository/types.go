// This file contains types that are used in the repository layer.
package repository

type RegisterUserInput struct {
	PhoneNumber string
	FullName    string
	Password    string
}

type UserData struct {
	Id          int
	PhoneNumber string
	FullName    string
	Password    string
}

type UpdateUserDataByIDInput struct {
	PhoneNumber string
	FullName    string
}
