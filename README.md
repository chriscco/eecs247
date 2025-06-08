# EECS 247 Project 2 

## Installation
Go dependencies installation: 

```bash
# producer/
go mod tidy 
``` 
Java dependencies: 
```bash
# project_247/ 
mvn clean compile 
``` 

gRPC Setup 
> Make sure ```protoc``` is installed on computer 

```bash
# producer/api/interface/server/proto/ 
protoc --go_out=. server.proto 
protoc --go-grpc_out=. server.proto 
``` 

## Instruction 
Run Java Server
```bash 
# project_247/src/main/java
javac Driver.java 
java Driver 
``` 
Run Go Client 
```bash 
# producer/ 
go run main.go 
``` 
Send a post request to ```localhost:8080/query``` with request: 
```json 
{
    "text": "hello world", 
}
```
