# go-postgreSQL-template

Go(GORM + Gin + Air) + PostgreSQL の API サーバー

## 使いかた

```
docker-compose up --build
```

postman とかで、`POST localhost:8080/memos` に以下の JSON を送ると、DB に保存される。

```json
{
  "content": "content"
}
```

localhost:8080/memos にアクセスすると、DB の中身が表示される。

## 参考記事
https://pontaro.net/1305/
