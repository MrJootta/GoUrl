# GoUrl

Simple too to make url shorter

### Concept

This application allows you to send a full url and get back a shorter version 

### Endpoints

**POST /create**

Endpoint to create new short url's it receives and returns json

_Request_

```
{   
    "url": "example.com"
}
```

_Response_

```
{
    "status": 201,
    "response": {
        "url": "http://localhost:2020/KaXTsk",
        "code": "KaXTsk"
    }
}
```

_Status Codes_

* 500 - When code is not generated or can't save in database
* 405 - When method is not GET
* 400 - When url param is not sent or url can't be parsed
* 201 - When short url is created

**GET /info/{code}**

Endpoint to consult number of visits a code got in the last 24h

_Response_

```
{
    "status": 200,
    "response": {
        "url": "http://localhost:2020/KaXTsk",
        "code": "KaXTsk",
        "number_of_visits": 3
    }
}
```

_Status Codes_

* 405 - When method is not GET
* 404 - When code is not found
* 200 - When returns visits count

**GET /{code}**

Endpoint to be redirected to the saved long url

_Status Codes_

* 405 - When method is not GET
* 404 - When code is not found
* 301 - When redirect happen

### Setup

To setup the application you just need to run the following commands

```
$ -> docker-compose up
$ -> go run main.go
```

### Current implementation

##### Services

- Mysql database
- Go Application

### Ideal Implementation

##### Services

- Mysql database
- Memcached
- Go Application
