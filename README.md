
Run tests locally
```
  docker pull postgres
  run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=test1234 -d postgres
  go test ./bank
```
Build docker image
```
  docker build -t bank-java .
  
```
Run app in Docker
with external postgres
```
  docker run --name bank-java -p 8080:8080 -e POSTGRES_HOST=${host} -d bank-javadocker build -t bank-java .
```
  or create postgres 
```
  docker-compose up -d --force-recreate
```


