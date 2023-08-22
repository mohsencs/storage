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
```bash
$ curl -X POST 'http://localhost:8080/api/promotion/upload' \
--header 'Content-Type: application/json' \
--form 'file=@"/home/user/promotions.csv"'
```

## Curl for get one promotion:
`$ curl -X GET 'http://localhost:8080/api/promotion/d018ef0b-dbd9-48f1-ac1a-eb4d90e57118' `
