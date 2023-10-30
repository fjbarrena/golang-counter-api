# Counter API using Golang

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

