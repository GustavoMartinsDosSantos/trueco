# TRUE-CO
Gerenciador de truco corporativo.
- Cadastro de jogadores
- Cadastro de duplas
- Partidas e apostas

Backend
- GOLANG
    - GINGONIC
    - GORM

# Install

Get golang packages
```
go mod tidy
```

Docker Build
```
docker-compose up --build
```
# Run
Start webservice
```
go run main.go
```
# Endpoints
- Player
    - GET /player/{id} Find player ID
    - GET /player Find all players
    - POST /player Create player 
    - POST /pair Create pair
