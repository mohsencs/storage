# storage
upload huge data and response to huge access

## How to run project:
```bash
#move to directory
$ cd workspace

# Clone into your workspace
$ git clone https://github.com/mohsencs/storage.git

#move to project
$ cd storage

$ docker-compose up
```

## Curl for uplad file:
`$ curl --location 'http://localhost:8080/api/promotion/upload' --header 'Content-Type: application/json'`

## Curl for get one promotion:
`$ curl --location 'http://localhost:8080/api/promotion/'`
