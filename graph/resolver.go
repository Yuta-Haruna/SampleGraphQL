package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"fmt"
	"log"

	"SampleGraphQL/graph/model"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Resolver 構造体はリゾルバに必要な依存性を注入するために使用します。
type Resolver struct {
	FirestoreClient *firestore.Client
}

// NewResolver は新しいリゾルバインスタンスを作成します。
func NewResolver(firestoreClient *firestore.Client) *Resolver {
	return &Resolver{
		FirestoreClient: firestoreClient,
	}
}

// InitFirestoreClient はFirestoreクライアントの初期化を行います。
func InitFirestoreClient(ctx context.Context) *firestore.Client {
	sa := option.WithCredentialsFile("../../credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase App: %v", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore Client: %v", err)
	}
	return client
}

// User はIDに基づいてユーザーを取得するリゾルバです。
func (r *Resolver) User(ctx context.Context, id string) (*model.User, error) {
	doc, err := r.FirestoreClient.Collection("user").Doc(id).Get(ctx)
	if err != nil {
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
