# Secrets Sharing Live Project

## Run the web server

```bash
go build -o secret-app
DATA_FILE_PATH=./data.json ./secret-app
```

## Test the web server

```bash
curl localhost:8080/
Hello from secret handler!%
```
