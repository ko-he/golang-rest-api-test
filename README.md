
## 概要
このリポジトリでは REST APIで開発されたAPIをフロントエンド、バックエンド間でスキーマ安全に保つためのテストとコード生成の方法をサンプルとして定義します。

## バックエンド
### テスト実行
```
go test ./...
```

## フロントエンド
### OpenAPI に定義されているスキーマを Typescript のモデルに置き換える
```
npx openapi-typescript etc/openapi/root.yml --output frontend/schema.ts
```

## おまけ
### OpenAPI スキーマから SwaggerUI をビルドしてブラウザで可視化
```
npx redoc-cli build etc/openapi/root.yml -o var/tmp/redoc/index.html
open var/tmp/redoc/index.html
```
