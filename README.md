# gate-link

## Development

```sh
# root
go run .

# gateway
go run . gateway

# gateway with auth
go run . gateway --auth

# version
go run . version

# generate
go run . generate
```

## Playground

```sh
# api users
curl -H "Authorization: Bearer your_token_here" http://localhost:8080/api/users

# api products
curl -H "Authorization: Bearer your_token_here" http://localhost:8080/api/products
```