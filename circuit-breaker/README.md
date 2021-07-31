# Demo with Circuite braker
* [hystrix-go](https://github.com/afex/hystrix-go/)

## Run with success

Start server
```
$go mod tidy
$go run main.go
```

Run client
```
$for i in $(seq 10); do curl -x '' localhost:8080 ;done
```

## Run with error
Start server
```
$go mod tidy
$export SERVER_ERROR=1 
$go run main.go
```

Run client
```
$for i in $(seq 10); do curl -x '' localhost:8080 ;done
```