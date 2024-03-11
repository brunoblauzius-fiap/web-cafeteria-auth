# web-cafeteria-auth

Repo com o Lambda que retorna token de autorização

Curl para testar:
```
curl --location 'https://au8b318jo9.execute-api.us-east-1.amazonaws.com/auth' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "primeirinho@gmail.com",
    "password": "Primeiro@123"
}' 
```

Desenho arquitetural (WIP)
