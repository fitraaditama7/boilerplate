# iseng

CRUD API Boilerplate

[Installation](#installation)

## Installation

### Prerequisites
- Golang(>=1.11) - Download and Install [Golang](https://golang.org/)
- Redis - Download and Install [Redis](https://redis.io/download)
- MySQL - Download and Install [MySQL](https://www.apachefriends.org/download.html)

### How To Use
- Download Dependency Golang

This command is useful for download dependency golang
```
    $ make download
```

This command is useful for migrate dummy data
```
    $ make migrate
```

- Running Server

This command is useful for running and generate swagger documentation
```
    $ make run
```

Curl for get list data
```
    $ curl --location --request GET 'localhost:4000/v1/dashboard/'
```

Curl for get detail data
```
    $ curl --location --request GET 'localhost:4000/v1/dashboard/dbe11835-0ae7-4d53-a6d7-f262d0b89603'
```

Curl for insert data
```
    $ curl --location --request PUT 'localhost:4000/v1/dashboard/dbe11835-0ae7-4d53-a6d7-f262d0b89603' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "first_name": "Fitra",
            "last_name": "Aditama",
            "phone_number": "0812345569",
            "email": "fitraaditama77@gmail.com",
            "role_id": "admin",
            "username": "fitraaditama77",
            "password": "akuganteng",
            "created_at": "fitra"
        }'
```

Curl for update data
```
    $ curl --location --request PUT 'localhost:4000/v1/dashboard/dbe11835-0ae7-4d53-a6d7-f262d0b89603' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "first_name": "Fitraa",
            "last_name": "Aditama",
            "phone_number": "0812345569",
            "email": "fitraaditama77@gmail.com",
            "role_id": "admin",
            "username": "fitraaditama77",
            "password": "akuganteng",
            "updated_at": "fitra"
}'
```

Curl for delete data
```
    $  curl --location --request DELETE 'localhost:4000/v1/dashboard/dbe11835-0ae7-4d53-a6d7-f262d0b89603' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "deleted_at": "fitra"
        }'
```

- Create Mock

This Command is useful for creating mock file from interface{}
```
    $ mockery --name=InterfaceName --output=cmd/mocks  --recursive
```
