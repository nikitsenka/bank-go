
Run tests locally
```
  docker pull postgres
  run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=test1234 -d postgres
  go test ./bank
```
Build docker image
```
  docker build --no-cache -t bank-go .
  
```
Run app in Docker
with external postgres
```
  docker run --name bank-go -p 8080:8080 -e POSTGRES_HOST=${host} -d bank-javadocker build -t bank-go .
```
  or create both postgres and bank-go containers and run 
```
  docker-compose up -d --force-recreate
```


