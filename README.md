
### OpenAPI スキーマから SwaggerUI をビルドしてブラウザで可視化
```
npx redoc-cli build etc/openapi/root.yml -o var/tmp/redoc/index.html
open var/tmp/redoc/index.html
```

### OpenAPI に定義されているスキーマを Typescript のモデルに置き換える
```
npx openapi-typescript etc/openapi/root.yml --output frontend/schema.ts
```
