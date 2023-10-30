# Counter API using Golang

## Building Docker image

```shell
docker build -t counter-api:latest .
```

## Running Docker image

```shell
docker run -it -e GOPORT=7777 -p 7777:7777 counter-api:latest
```

## Running endpoint

```shell
curl -X 'GET' 'http://localhost:7777/status'
{"counter":1,"status":true,"lastChanged":"2023-10-30T14:50:07.021844932Z"}

curl -X 'GET' 'http://localhost:7777/status'
{"counter":2,"status":true,"lastChanged":"2023-10-30T14:50:30.189676493Z"}
```

