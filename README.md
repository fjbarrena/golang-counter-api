# Counter API using Golang

## Running project locally

```shell
$ go run api.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /status                   --> counter-api/controllers.UpdateStatus (4 handlers)
[GIN-debug] GET    /shutdown                 --> counter-api/controllers.Shutdown (4 handlers)
[GIN-debug] GET    /poweron                  --> counter-api/controllers.PowerOn (4 handlers)
[GIN-debug] Listening and serving HTTP on :40000
```

To change the default port use `GOPORT` env variable

```shell
$ GOPORT=7777 go run api.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /status                   --> counter-api/controllers.UpdateStatus (4 handlers)
[GIN-debug] GET    /shutdown                 --> counter-api/controllers.Shutdown (4 handlers)
[GIN-debug] GET    /poweron                  --> counter-api/controllers.PowerOn (4 handlers)
[GIN-debug] Listening and serving HTTP on :7777
```

## Running tests

```shell
$ go test -cover ./...
?       counter-api     [no test files]
?       counter-api/models      [no test files]
?       counter-api/controllers [no test files]
ok      counter-api/services    0.001s  coverage: 100.0% of statements
```

## Building Docker image

```shell
docker build -t counter-api:latest .
```

## Running Docker image

```shell
docker run -it -e GOPORT=7777 -p 7777:7777 counter-api:latest
```

## Running endpoint

> Counter increments on each request, but lastChanged doesn't change until the status changes

```shell
curl -X 'GET' 'http://localhost:7777/status'
{"counter":1,"status":true,"lastChanged":"2023-10-30T14:50:07.021844932Z"}

curl -X 'GET' 'http://localhost:7777/status'
{"counter":2,"status":true,"lastChanged":"2023-10-30T14:50:07.021844932Z"}
```

> Changing status to false (/shutdown) provokes a change as well at lastChanged, but not at counter

```shell
curl -X 'GET' 'http://localhost:7777/shutdown'
{"counter":2,"status":false,"lastChanged":"2023-10-30T17:11:35.419677326Z"}
```

> Same changing status to true (/poweron) provokes a change as well at lastChanged, but not at counter

```shell
curl -X 'GET' 'http://localhost:7777/poweron'
{"counter":2,"status":true,"lastChanged":"2023-10-30T17:11:39.131874167Z"}
```

