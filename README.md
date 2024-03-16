# web-cafeteria-auth

Repo com o Lambda que retorna token de autorização

Curl para testar:
```
curl --location 'https://7cylcgukg9.execute-api.us-east-1.amazonaws.com/authentication' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "john@doe.com",
    "password": "John@doe123"
}'

```

Video: https://www.loom.com/share/e1b61d78c2424c0ebcfbea44d80fb9c4
