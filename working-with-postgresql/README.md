# Manage depednency in project
* Global variable
* Depednency Injection
* Wrapper
* Request context

### Step to run

Start database server
```
$docker-compose up -d
$docker-compose ps
```


Start API Server
```
$go mod tidy
$go run cmd/main.go
```