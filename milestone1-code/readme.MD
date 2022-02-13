# Secrets Sharing Live Project

## Run the web server

```bash
go build -o secret-app
DATA_FILE_PATH=./data.json ./secret-app
```

## Test the web server

```bash
# GET
curl localhost:8080/
GET/secret handler

# POST
curl -X "POST" localhost:8080/
POST/secret handler
```
