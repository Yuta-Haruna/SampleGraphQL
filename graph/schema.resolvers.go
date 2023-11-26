package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"SampleGraphQL/graph/model"
	"context"
	"fmt"
	"log"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	doc, err := r.FirestoreClient.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		// Firestoreからの取得に失敗した場合のエラーハンドリング
		log.Printf("Error fetching user with ID %s: %v", id, err)
		return nil, fmt.Errorf("error fetching user: %v", err)
	}

	var user model.User
	if err := doc.DataTo(&user); err != nil {
		// データ変換に失敗した場合のエラーハンドリング
		return nil, err
	}
	// FirestoreのドキュメントIDをユーザーモデルのIDフィールドに設定
	user.ID = doc.Ref.ID
	return &user, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
