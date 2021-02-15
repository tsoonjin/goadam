module github.com/tsoonjin/goadam

go 1.14

require (
	github.com/Jeffail/gabs/v2 v2.6.0
	github.com/joho/godotenv v1.3.0
)

replace github.com/tsoonjin/goadam/internal/service => ./internal/service
