// This file contains types that are used in the repository layer.
package repository

import "github.com/google/uuid"

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type UserModel struct {
	ID          uuid.UUID
	PhoneNumber string
	Password    string
	FullName    string
}
