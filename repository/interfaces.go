// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	RegisterUser(ctx context.Context, input RegisterUserInput) (output int, err error)
	GetUserDataByPhoneNumber(ctx context.Context, input string) (output UserData, err error)
	GetUserDataByID(ctx context.Context, input int) (output UserData, err error)
	UpdateUserDataByID(ctx context.Context, input UserData) (err error)
	UpdateLoginActivity(ctx context.Context, input int) (err error)
}
