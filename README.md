#Quick Start
##Build

```
  go build -o bin/bank ./bank/ 
```

##Run Application locally

```
  export PORT=8000
  chmod 500 ./bin/bank
  ./bin/bank
```

##Run DB locally

```
  docker build --no-cache -t bank-postgres ./docker/db/
  docker run --name bank-postgres -p 5432:5432 -e POSTGRES_PASSWORD=test1234 -d bank-postgres
```

##Test

```
  curl -X POST http://localhost:8000/client/new/100 
```

```Output
  {"id":1,"name":"","email":"","phone":""}
```

