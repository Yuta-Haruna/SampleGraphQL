package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "cloud.google.com/go/firestore"

type Resolver struct {
	client *firestore.Client
}
