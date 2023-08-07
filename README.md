# SampleGraphQL
【未完成】
状況：GraphQL PlayGraundの起動まで確認

　　　クエリは失敗

## 背景
別アプリにて、GraphQLサーバの構築が上手くできなかったため、技術検証用のプロジェクト

## 実行方法

①プロジェクトをクローン
 ```sh
https://github.com/Yuta-Haruna/SampleGraphQL.git
 ```

②ターミナルにてクローン先のディレクトリに移動し、下記コマンドを入力
 ```sh
 go run server.go 
 ```

③下記URLをブラウザに入力する
 ```sh
http://localhost:8080/graphql
 ```

## 最終目標
GraphQLサーバを起動し、下記クエリにてFireStore内の情報を表示する

 ```sh
query {
    breads {
        id
        name
        createdAt
    }
}
 ```

 ```sh
query {
    breads (id='FireStore内の取得したいID'){
        id
        name
        createdAt
    }
}
 ```
