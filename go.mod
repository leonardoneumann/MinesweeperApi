module github.com/leonardoneumann/minesweeperapi

go 1.12

require (
	github.com/leonardoneumann/minesweeperapi/api v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.4.2 // indirect
)

replace github.com/leonardoneumann/minesweeperapi/api => ./internal
