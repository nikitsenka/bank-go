```
  docker pull postgres
  run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=test1234 -d postgres
  go test ./bank
```

