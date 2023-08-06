package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/graphql-go/graphql"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// パンの情報を表す構造体
type Bread struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func main() {
	// Firestoreクライアントの初期化
	ctx := context.Background()
	opt := option.WithCredentialsFile("credentials.json")
	client, err := firestore.NewClient(ctx, "your-project-id", opt)
	if err != nil {
		log.Fatalf("Firestoreクライアントの初期化に失敗しました: %v", err)
	}
	defer client.Close()

	// GraphQLスキーマの定義
	var breadType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Bread",
			Fields: graphql.Fields{
				"id":        &graphql.Field{Type: graphql.String},
				"name":      &graphql.Field{Type: graphql.String},
				"createdAt": &graphql.Field{Type: graphql.String},
			},
		},
	)

	fields := graphql.Fields{
		"breads": &graphql.Field{
			Type: graphql.NewList(breadType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Firestoreからデータを取得
				iter := client.Collection("breads").Documents(ctx)
				defer iter.Stop()

				var breads []Bread
				for {
					doc, err := iter.Next()
					if err == iterator.Done {
						break
					}
					if err != nil {
						log.Printf("Firestoreデータ取得エラー: %v", err)
						return nil, err
					}

					var bread Bread
					if err := doc.DataTo(&bread); err != nil {
						log.Printf("Firestoreデータ変換エラー: %v", err)
						return nil, err
					}
					breads = append(breads, bread)
				}

				return breads, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("スキーマの作成に失敗しました: %v", err)
	}

	// GraphQLハンドラの作成
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		if len(result.Errors) > 0 {
			log.Printf("クエリエラー: %v", result.Errors)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", result.Data)
	})

	// サーバの起動
	port := ":4000"
	log.Printf("サーバが http://localhost%s/graphql で起動しました。", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
