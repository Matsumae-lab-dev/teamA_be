# go-postgreSQL-template

Go(GORM + Gin + Air) + PostgreSQL の API サーバー

## 使いかた

```
docker-compose up -d
```

postman とかで、`POST localhost:8080/todos` に以下の JSON を送ると、DB に保存される。

```json
{
    "Title": "あああ",
    "Description": "こんにちは",
    "Category": "category3",
    "Deadline": "2006-01-02T00:00:00Z",
    "State": false
}
```

localhost:8080/todos にアクセスすると、DB の中身が表示される。

## 参考記事
https://pontaro.net/1305/
