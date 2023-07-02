# Shell Command Executor

Golang based binary to execute shell commands.

## Library Used :
1. Golang fiber
2. Standard os library


Users can execute use this binary to execute shell commands.

You would get the response as :-
````
Output
````

## How to use it
- Build and execute the binary.
- Use postMan or Curl command to send POST request

```bash
curl -X POST -H "Content-Type: application/json" -d '{"command": "ls"}' http://localhost:3000/api/cmd
```

How to Deploy it :
- You can use docker-compose up -d to deploy it using docker and use endpoint as 127.0.0.1:3000
- You can use the local host to deploy it locally

Sample Request :
- API requests can be sent using POSTMAN or CURL
- Once the path is set use
 
```bash
Example :
Request : curl -H "Content-Type: application/json" -X POST -d '{"command":"ls"}' http://localhost:3000/api/cmd
Response : {"Output":"docker-compose.yaml, Dockerfile, docs, go.mod, go.sum, main.go, models, README.md, routers, shell-command-executor.exe"}
```


