# Develop RESTful API with Gin framework
* Better structure
* Dependency Injection
* Middleware
* Observability
  * Application metrics
  * Distributed tracing
  * Logging

### Step to run
```
$go mod tidy
$go run cmd/main.go
```
List of APIs
* GET http://localhost:8080/api/v1/user
* GET http://localhost:8080/api/v1user/<id>
* GET http://localhost:8080/metrics
* POST http://localhost:8080/api/v1/login
* POST http://localhost:8080/api/v1/signup
