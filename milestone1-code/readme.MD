# Secrets Sharing Live Project

## Run the web server

```bash
go build -o secret-app
DATA_FILE_PATH=./data.json ./secret-app
```

## Test the web server

```bash
# GET
curl localhost:8080/ca63dc2acc2522785933852d4ae3b181
{"data":"my$sdfsdfuper$secret"}

# POST
curl -X "POST" \
  -H "Content-Type: application/json" \
  -d '{"plain_text":"my$super$secret"}' \
  localhost:8080

POST/secret handler
```
