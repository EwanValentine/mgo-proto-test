docker run -d -p 8500:8500 consul || true
docker run -d -p 27017:27017 mongo || true
go run main.go handlers.go repository.go --server_address=0.0.0.0:50051
