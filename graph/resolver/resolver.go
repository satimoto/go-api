package resolver

import "github.com/satimoto/go-datastore/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository db.Repository
}
