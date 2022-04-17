### Golang Rest Api with Docker
- Golang
- Gin
- Test
- Mvc Design 
- Restapi

### let's start
for build
```console
docker-compose up --build
```
start
```console
docker-compose up
```
stop 

```console
docker-compose down
```
for run unit-test
```console

go test ./tests/ -v
```
#### We replicated the application from 2 separate ports via docker-compose
App is ready.
<br>
<br>
**why two ports ?** 
- for high traffic (desire state)
 
[localhost:8888](http://localhost:8888)
<br>
[localhost:8889](http://localhost:8889)
