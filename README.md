# storage
upload huge data and response to huge access

for run docker:
$ docker-compose up

for request to upload promotion file use this curl:
$ curl --location 'http://localhost:8080/api/promotion/upload' --header 'Content-Type: application/json'

for access to promotion by id, use this curl:
$ curl --location 'http://localhost:8080/api/promotion/'
