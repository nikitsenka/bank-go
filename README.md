
Run tests locally
```
  docker pull postgres
  run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=test1234 -d postgres
  go test ./bank
```

Run app in Docker
```
  docker build --no-cache -t bank-go .
  docker-compose up -d --force-recreate
```


