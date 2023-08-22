# storage
upload huge data and response to huge access

## How To Run This Project
```bash
#move to directory
$ cd workspace

# Clone into your workspace
$ git clone https://github.com/mohsencs/storage.git

#move to project
$ cd storage

$ docker-compose up
```

## for request to upload promotion file use this curl:
`$ curl --location 'http://localhost:8080/api/promotion/upload' --header 'Content-Type: application/json'`

## for access to promotion by id, use this curl:
`$ curl --location 'http://localhost:8080/api/promotion/'`
