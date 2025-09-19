package pkg

import "errors"

var (
	ErrNotFound           = errors.New("your requested Item was not found")
	ErrClerkUserNotSynced = errors.New("user for the given clerk_id does not exist")
	ErrUserIdNull         = errors.New("user_id field cannot be null")
	ErrLoadingENV         = errors.New("error loading .env file")
	ErrConnectionDB       = errors.New("cannot establish connection to the database")
	ErrMigratingDB        = errors.New("auto migration to DB failed")
)
