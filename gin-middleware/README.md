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
* http://localhost:8080/api/v1/user
* http://localhost:8080/api/v1user/<id>
* http://localhost:8080/metrics
